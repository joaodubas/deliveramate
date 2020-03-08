# go parameters
GOMOD_PATH=${GOPATH}/pkg/mod
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMG) get

# binary names
BUILD_PATH=build
LOADER_NAME=$(BUILD_PATH)/deliveramate-loader
LOADER_UNIX=$(LOADER_NAME)-unix
LOADER_WIN=$(LOADER_NAME).exe
LOADER_OSX=$(LOADER_NAME)-darwin
SERVER_NAME=$(BUILD_PATH)/deliveramate-server
SERVER_UNIX=$(SERVER_NAME)-unix
SERVER_WIN=$(SERVER_NAME).exe
SERVER_OSX=$(SERVER_NAME)-darwin
CLIENT_NAME=$(BUILD_PATH)/deliveramate-client
CLIENT_UNIX=$(CLIENT_NAME)-unix
CLIENT_WIN=$(CLIENT_NAME).exe
CLIENT_OSX=$(CLIENT_NAME)-darwin

.DEFAULT_GOAL := help

.PHONY: setup
setup:  ## install application packages and copy proper protobuf files
	$(GOCMD) mod download
	cp --recursive /opt/src/protoc-3.11.4-linux-x86_64/include/google third_party
	cp --recursive $(GOMOD_PATH)/github.com/grpc-ecosystem/grpc-gateway\@v1.14.1/third_party/googleapis/google third_party
	mkdir -p third_party/protoc-gen-swagger/options
	cp $(GOMOD_PATH)/github.com/grpc-ecosystem/grpc-gateway\@v1.14.1/protoc-gen-swagger/options/annotations.proto third_party/protoc-gen-swagger/options
	cp $(GOMOD_PATH)/github.com/grpc-ecosystem/grpc-gateway\@v1.14.1/protoc-gen-swagger/options/openapiv2.proto third_party/protoc-gen-swagger/options

.PHONY: protoc-gen
protoc-gen: protoc-gen-grpc protoc-gen-gateway protoc-gen-openapi  ## generate golang code based in protobuf files

.PHONY: protoc-gen-grpc
protoc-gen-grpc:  ## generate golang grpc code based in protobuf files
	mkdir -p pkg/http/grpc/v1
	protoc -I=api/proto/v1 -I=third_party --go_out=plugins=grpc:pkg/http/grpc/v1 api/proto/v1/partner-service.proto

.PHONY: protoc-gen-gateway
protoc-gen-gateway:  ## generate golang http gateway code based in protobuf files
	mkdir -p pkg/http/grpc/v1
	protoc -I=api/proto/v1 -I=third_party --go_out=plugins=grpc:pkg/http/grpc/v1 --grpc-gateway_out=logtostderr=true:pkg/http/grpc/v1 api/proto/v1/partner-service.proto

.PHONY: protoc-gen-openapi
protoc-gen-openapi:  ## generate openapi spec based in protobuf files
	mkdir -p api/openapi/v1
	protoc -I=api/proto/v1 -I=third_party --swagger_out=logtostderr=true:api/openapi/v1 api/proto/v1/partner-service.proto

.PHONY: test
test:  ## execute tests
	$(GOTEST) -v ./...

.PHONY: coverage
coverage:  ## execute coverage report
	$(GOTEST) -v -cover -coverprofile=coverage.out ./...

build: build-linux build-windows build-osx ## build commands for all os

build-linux:  ## build commands for linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LOADER_UNIX) -v ./cmd/loader/...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(SERVER_UNIX) -v ./cmd/server/...

build-windows:  ## build commands for windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(LOADER_WIN) -v ./cmd/loader/...
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(SERVER_WIN) -v ./cmd/server/...

build-osx:  ## build commands for osx
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(LOADER_OSX) -v ./cmd/loader/...
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(SERVER_OSX) -v ./cmd/server/...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
