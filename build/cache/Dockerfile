## Building Image
FROM golang:latest as builder

#ROOT := github.com/caicloud/nirvana-practice
ARG WORKSPACE=/practice
ARG PROJECT_ROOT=github.com/caicloud/nirvana-practice
ARG VERSION="v0.1.0"
ARG COMMIT="123"
ARG BRANCH=master

WORKDIR $WORKSPACE/src/$PROJECT_ROOT

COPY cmd/cache cmd/cache
COPY pkg pkg
COPY vendor vendor
COPY Makefile Makefile

ENV GOPATH=$GOPATH:$WORKSPACE

RUN    	CGO_ENABLED=0 GOOS=linux GOARCH=amd64						\
    	go build -v -o bin/practice-cache                          \
    	  -ldflags "-s -w                                              \
    	            -X  $PROJECT_ROOT/pkg/info.version=$VERSION         \
    	            -X  $PROJECT_ROOT/pkg/info.commit=$COMMIT           \
    	            -X  $PROJECT_ROOT/pkg/info.branch=$BRANCH"      \
    	  ./cmd/cache

## Running Image
FROM debian:jessie

ARG PORT=8081

ARG WORKSPACE=/practice
ARG PROJECT_ROOT=github.com/caicloud/nirvana-practice

COPY --from=0 $WORKSPACE/src/$PROJECT_ROOT/bin/practice-cache .

ENV PORT=${PORT}

CMD ["/bin/bash", "-c", "/practice-cache -p=${PORT}"]

EXPOSE ${PORT}
