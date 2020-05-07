#LICENSE_DIR=./licenses/
#BUILD_DIR=./build
#DOCKER_DIR=./docker/
SHELL := /bin/zsh
GO_VERSION=`cat GO_VERSION`
#DOCKER_BUILD_IMAGE=gotify/build
#DOCKER_WORKDIR=/proj
#DOCKER_RUN=docker run --rm -v "$$PWD/.:${DOCKER_WORKDIR}" -v "`go env GOPATH`/pkg/mod/.:/go/pkg/mod:ro" -w ${DOCKER_WORKDIR}
#DOCKER_GO_BUILD=go build -mod=readonly -a -installsuffix cgo -ldflags "$$LD_FLAGS"


default:
	@go build -ldflags '-linkmode external -w -extldflags "-static"' -o release/gin-wechat

#build:go build -ldflags '-linkmode external -w -extldflags "-static"' -o release/gin-wechat

.PHONY: build
