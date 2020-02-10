.DEFAULT_GOAL := help

.PHONY: protoc-gen
protoc-gen:  ## generate golang code based in protobuf files
	mkdir -p pkg/http/grpc/v1
	protoc -I=api/proto/v1 --go_out=plugins=grpc:pkg/http/grpc/v1 api/proto/v1/partner-service.proto

.PHONY: test
test:  ## execute tests
	go test -v ./...

.PHONY: coverage
coverage:  ## execute coverage report
	go test -v -cover -coverprofile=coverage.out ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
