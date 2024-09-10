## options
# based on https://tech.davis-hansson.com/p/make/
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
.DEFAULT_GOAL := help
.DELETE_ON_ERROR:

## variable
PROTO_ROOT="pb"
PROTO_STATIC_PATH = "${PROTO_ROOT}/static"
PROTO_DYNAMIC_PATH = "${PROTO_ROOT}/dynamic"
PROTO_DESCRIPTOR_OUTPATH = "artefacts"

## formula
.PHONY: help
help:  ## print help message
	@grep -E '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## checks
.PHONY: test
test: ## run unit tests
	go test -v ./...

## proto

.PHONY: proto-setup
proto-go-setup: ## download google.golang.org/protobuf & protoc-gen-go plugin
	go get google.golang.org/protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: protogen
protogen: ## compile all proto files in proto/* into go files
	protoc --proto_path=${PROTO_STATIC_PATH} --go_out=. ${PROTO_STATIC_PATH}/*

.PHONY: protoc-descriptor
protoc-descriptor: ## compile all proto files to its binary descriptor format
	protoc --proto_path=${PROTO_DYNAMIC_PATH} --include_imports -o ${PROTO_DESCRIPTOR_OUTPATH}/bin.desc ${PROTO_DYNAMIC_PATH}/* 

## service

.PHONY: dev
dev: ## starts a dev server
	go run ./cmd/server.go