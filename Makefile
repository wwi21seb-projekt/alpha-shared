GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	@protoc --proto_path=./proto/post --go_out=./proto/post --go_opt=paths=source_relative \
		--go-grpc_out=./proto/post --go-grpc_opt=paths=source_relative proto/post/*.proto
	@protoc --proto_path=./proto/health --go_out=./proto/health --go_opt=paths=source_relative \
		--go-grpc_out=./proto/health --go-grpc_opt=paths=source_relative proto/health/*.proto