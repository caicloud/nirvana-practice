
# Target binaries. You can build multiple binaries for a single project.
TARGETS ?= cache server

# Container registry,
REGISTRY ?= cargo.caicloudprivatetest.com/caicloud

# Current version of the project.
VERSION ?= $(COMMIT)
# Git commit sha.
COMMIT := $(shell git rev-parse HEAD)
# Git branch name,
BRANCH := $(shell git branch | grep \* | cut -d ' ' -f2))

IMAGE_PREFIX := practice-

# A list of all packages.
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /test)

# This repo's root import path (under GOPATH).
ROOT := github.com/caicloud/nirvana-practice
# Project main package location (can be multiple ones).
CMD_DIR := ./cmd
# Project output directory.
OUTPUT_DIR := ./bin
# Build direcotory.
BUILD_DIR := ./build

# All targets.
.PHONY: test build build-local build-linux container push

build: build-local

test:
	go test $(PKGS)

build-local:
	@for target in $(TARGETS); do                                                      \
	  CGO_ENABLED=0																	   \
	  go build -v -o $(OUTPUT_DIR)/$(IMAGE_PREFIX)$${target}						   \
	    -ldflags "-s -w										                      	   \
	              -X  $(ROOT)/pkg/info.version=$(VERSION)							   \
	              -X  $(ROOT)/pkg/info.commit=$(COMMIT)								   \
	              -X  $(ROOT)/pkg/info.branch=$(BRANCH)"							   \
	    $(CMD_DIR)/$${target};                                                         \
	done

build-linux:
	@for target in $(TARGETS); do                                                      \
	  CGO_ENABLED=0 GOOS=linux GOARCH=amd64											   \
	  go build -v -o $(OUTPUT_DIR)/$(IMAGE_PREFIX)$${target}						   \
	    -ldflags "-s -w										                      	   \
	              -X  $(ROOT)/pkg/info.version=$(VERSION)							   \
	              -X  $(ROOT)/pkg/info.commit=$(COMMIT)								   \
	              -X  $(ROOT)/pkg/info.branch=$(BRANCH)"							   \
	    $(CMD_DIR)/$${target};                                                         \
	done

container: test build-linux
	@for target in $(TARGETS); do                                                      \
	  image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                  \
	  docker build -t $(REGISTRY)/$${image}:$(VERSION)                                 \
	    -f $(BUILD_DIR)/$${target}/Dockerfile .;                                       \
	done

push: container
	@for target in $(TARGETS); do                                                      \
	  image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                  \
	  docker push $(REGISTRY)/$${image}:$(VERSION);                                    \
	done
