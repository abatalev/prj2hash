#!/bin/sh
CDIR=$(pwd)
if [ ! -d "build" ]; then
    mkdir build
fi 
if [ ! -f "build/gototcov" ]; then
    cd build
    git clone https://github.com/jonaz/gototcov gototcov.git
    cd gototcov.git
    go get golang.org/x/tools/cover
    go build .
    cp gototcov.git ../gototcov
    cd ${CDIR}/build
    rm -f -R gototcov.git
    echo "### done build tools"
fi

cd ${CDIR}
echo "### calc coverage"
go test -coverprofile=coverage.out .
if [ "$?" != "0" ]; then
    echo "### aborted"
    exit 1
fi

echo "### total coverage"
./build/gototcov -f coverage.out -limit 60
if [ "$?" != "0" ]; then
    echo "### open browser"
    go tool cover -html=coverage.out
    echo "### aborted"
    exit 1
fi

echo "### build application"
GIT_HASH=$(git rev-list -1 HEAD)
if [ -f "./prj2hash" ]; then
    P2H_HASH=$(./prj2hash)
fi
go build -ldflags "-X main.gitHash=${GIT_HASH} -X main.p2hHash=${P2H_HASH}" .
echo "### run new version info"
./prj2hash -version
echo "### the end"
