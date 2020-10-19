BIN := bin
LOGDIR := logs
PWD := $(shell pwd)

PROTOC_DOCKERFILE ?= protoc.Dockerfile
PROTOC_IMAGE ?= mukgo_protoc

API_SERVER_NAME ?= mukgo_api
DB_SERVER_NAME ?= mukgo_db
LOG_SERVER_NAME ?= mukgo_log

.DEFAULT_GOAL := help
.PHONY: help api db protocol log clean _build_protoc

server: api db log ## Build all server

api: ## Build api server
	go build -o $(BIN)/$(API_SERVER_NAME) ./server/cmd/mukgo-api

db: ## Build db server
	go build -o $(BIN)/$(DB_SERVER_NAME) ./server/cmd/mukgo-db

log: ## Build log server
	go build -o $(BIN)/$(LOG_SERVER_NAME) ./server/cmd/mukgo-log

protocol: _build_protoc ## Compile proto file into dart, go files
	docker run --rm -it \
		-v $(PWD)/proto:/proto \
		-v $(PWD)/client/lib:/protocol_dart \
		-v $(PWD)/server/api:/protocol_go \
		$(PROTOC_IMAGE)

clean: ## Clean binaries
	rm -f $(BIN)/$(API_SERVER_NAME)
	rm -f $(BIN)/$(DB_SERVER_NAME)
	rm -f $(BIN)/$(LOG_SERVER_NAME)
	rm -rf $(LOGDIR)

_build_protoc:
	docker build \
		--network=host \
		-f $(PROTOC_DOCKERFILE) \
		-t $(PROTOC_IMAGE) \
		.

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'