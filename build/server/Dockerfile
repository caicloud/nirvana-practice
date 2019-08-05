## Building Image
FROM golang:latest as builder

#ROOT := github.com/caicloud/nirvana-practice
ARG WORKSPACE=/practice
ARG PROJECT_ROOT=github.com/caicloud/nirvana-practice
ARG VERSION="v0.1.0"
ARG COMMIT="123"
ARG BRANCH=master

WORKDIR $WORKSPACE/src/$PROJECT_ROOT

COPY cmd/server cmd/server
COPY pkg pkg
COPY vendor vendor
COPY Makefile Makefile

ENV GOPATH=$GOPATH:$WORKSPACE

RUN    	CGO_ENABLED=0 GOOS=linux GOARCH=amd64						\
    	go build -v -o bin/practice-server                          \
    	  -ldflags "-s -w                                              \
    	            -X  $PROJECT_ROOT/pkg/info.version=$VERSION         \
    	            -X  $PROJECT_ROOT/pkg/info.commit=$COMMIT           \
    	            -X  $PROJECT_ROOT/pkg/info.branch=$BRANCH"      \
    	  ./cmd/server


## Running Image
FROM debian:jessie as prod

ARG WORKSPACE=/practice
ARG PROJECT_ROOT=github.com/caicloud/nirvana-practice

COPY --from=0 $WORKSPACE/src/$PROJECT_ROOT/bin/practice-server .

ARG LISTENPORT=8080

ARG CACHEADDRESS=localhost
ARG CACHEPORT=8081

ARG WORKSPACE=/practice

ENV LISTENPORT=${LISTENPORT}
ENV CACHEADDRESS=${CACHEADDRESS}
ENV CACHEPORT=${CACHEPORT}

CMD ["/bin/bash", "-c", "/practice-server -p=${LISTENPORT} --cache-endpoint=${CACHEADDRESS} --cache-port=${CACHEPORT}"]

EXPOSE ${LISTENPORT}




