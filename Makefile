GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=./proto/post --micro_out=./proto/post --micro_opt=paths=source_relative --go_out=:./proto/post --go_opt=paths=source_relative  proto/post/*.proto
	@protoc --proto_path=./proto/health --micro_out=./proto/health --micro_opt=paths=source_relative --go_out=:./proto/health --go_opt=paths=source_relative  proto/health/*.proto