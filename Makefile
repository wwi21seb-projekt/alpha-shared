GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	@protoc --proto_path=./proto --go_out=./proto --go_opt=paths=source_relative \
		--go-grpc_out=./proto --go-grpc_opt=paths=source_relative proto/server_alpha/*/*.proto