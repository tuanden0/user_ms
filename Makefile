STATICLDFLAGS := -ldflags '-s -extldflags "-static"'
PWD = $(shell pwd)

.PHONY: install update build

install:
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.7
	go install github.com/envoyproxy/protoc-gen-validate@v0.4.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

build:
	go install ./tools/build

update:
	go run ./tools/build protoc