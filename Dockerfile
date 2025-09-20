FROM registry.access.redhat.com/ubi8/go-toolset:1.17.12-11 AS build
COPY ./vendor $APP_ROOT/src/vendor
COPY go.* $APP_ROOT/src
COPY build.sh $APP_ROOT/src
COPY internal/ $APP_ROOT/src/internal/
COPY cmd $APP_ROOT/src/cmd/

COPY examples/ $APP_ROOT/src/examples/


WORKDIR $APP_ROOT/src
RUN mkdir $APP_ROOT/tmp
ENV TMPDIR=$APP_ROOT/tmp
RUN go build -o prj2hash cmd/prj2hash/main.go
#RUN chown root:root /opt/app-root/src/prj2hash

FROM registry.access.redhat.com/ubi8/ubi-micro:8.6-526
USER 1001
COPY --chown=1001 --from=build /opt/app-root/src/prj2hash .
CMD [ "./prj2hash" ]