#!/bin/sh
CDIR=$(pwd)
if [ ! -d "build" ]; then
    mkdir build
fi 
if [ ! -f "build/gototcov" ]; then
    cd build
    git clone https://github.com/jonaz/gototcov gototcov.git
    cd gototcov.git
    go get: added golang.org/x/tools v0.1.8
    go build .
    cp gototcov.git ../gototcov
    cd ${CDIR}/build
    rm -f -R gototcov.git
    echo "### done build tools"
fi

cd ${CDIR}
echo "### coverage"
go test -coverprofile=coverage.out .
# go tool cover -html=coverage.out
echo "### total coverage"
./build/gototcov -f coverage.out -limit 60
if [ "$?" != "0" ]; then
    echo "### aborted"
    exit 1
fi

echo "### build application"
go build .
echo "### the end"
