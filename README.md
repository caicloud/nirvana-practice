# Nirvana Practice Project

<img align="right" width="225px" src="https://user-images.githubusercontent.com/2191361/35839723-e9e5cdfa-0b2c-11e8-853a-8d3870f9e7ac.png">

As a HTTP server framework, [Nirvana](https://github.com/caicloud/nirvana) is robust but can also be quite complicated. This project let you practice writing Restful APIs with Nirvana.

## Prerequisites

 - [Golang](https://golang.org/dl/) - the Go language
 - [dep](https://github.com/golang/dep) - the Go dependency management tool used by Nirvana
 
That is. You'll use dep to install the required Go packages.

## Getting Started

The project is already runnable as it is. Use dep to install the dependencies and the Makefile to build it.

```
$ dep ensure -v

$ VERSION=v0.1.0 make build
```

Check out the help message and version information.

```
 $ bin/practice-server -h
Usage of bin/practice-server:
  -p, --port uint16   the HTTP port used by the server (default 8080)
  -v, --version       show version info
pflag: help requested

$ bin/nirvana-practice -v
nirvana-practice, version v0.1.0 (branch: master), revision: ab9132625bd5e2b9c8d3b2ff23c423cc9c02e9ec)
```

Now run the server and try the APIs.

```
$ bin/practice-server -p 8080

$ curl localhost:8080/api/v1alpha1/products
{"reason":"practice:NotImplemented","message":"requested feature is not implemented"}

$ curl localhost:8080/api/v1alpha1/unknown
{"reason":"Nirvana:Router:RouterNotFound","message":"can't find router"}
```

## Do It Yourself

The product API is already defined in `pkg/apis/v1alpha1/products.go` for you, alone with the incomplete `GET` and `LIST` API descriptors and handlers. Try to finish the `GET` and `LIST` APIs, and write `POST`, `LIST`, `PUT`, and `DELETE` from scratch. 

Use GitHub issue to ask questions and, once finished, submit a PR to this repo to show off your result. 

## Containerize Your App

Once your code is finished and tested, you might want to containerize your app so that you can deploy it practically. The Dockerfile is created for you under the `/build` directory. Try to finish it run `make container` to build the container images. 

Check out [the official Docker Getting-Started guide](https://docs.docker.com/get-started/) if you are not already familiar with Docker and containers. 
