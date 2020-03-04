# go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMG) get

# binary names
LOADER_NAME=deliveramate-loader
LOADER_UNIX=$(LOADER_NAME)-unix
LOADER_WIN=$(LOADER_NAME).exe
LOADER_OSX=$(LOADER_NAME)-darwin
SERVER_NAME=deliveramate-server
SERVER_UNIX=$(SERVER_NAME)-unix
SERVER_WIN=$(SERVER_NAME).exe
SERVER_OSX=$(SERVER_NAME)-darwin
CLIENT_NAME=deliveramate-client
CLIENT_UNIX=$(CLIENT_NAME)-unix
CLIENT_WIN=$(CLIENT_NAME).exe
CLIENT_OSX=$(CLIENT_NAME)-darwin

.DEFAULT_GOAL := help

.PHONY: protoc-gen
protoc-gen:  ## generate golang code based in protobuf files
	mkdir -p pkg/http/grpc/v1
	protoc -I=api/proto/v1 --go_out=plugins=grpc:pkg/http/grpc/v1 api/proto/v1/partner-service.proto

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
