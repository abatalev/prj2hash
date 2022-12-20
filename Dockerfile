FROM registry.access.redhat.com/ubi8/go-toolset:1.17.12-11 as build
COPY ./vendor $APP_ROOT/src/vendor
COPY ./main.go $APP_ROOT/src/main.go
COPY ./main_test.go $APP_ROOT/src/main_test.go
COPY ./go.mod $APP_ROOT/src/go.mod

WORKDIR $APP_ROOT/src
RUN mkdir $APP_ROOT/tmp
ENV TMPDIR=$APP_ROOT/tmp
RUN go build .
#RUN chown root:root /opt/app-root/src/prj2hash

FROM registry.access.redhat.com/ubi8/ubi-micro:8.6-526
USER 1001
COPY --chown=1001 --from=build /opt/app-root/src/prj2hash .
CMD ./prj2hash