BIN := bin
API_SERVER_NAME ?= mukgo_api
LOG_SERVER_NAME ?= mukgo_log

.DEFAULT_GOAL := help
.PHONY: help api log clean

server: api log ## Build all server

api: ## Build api server
	go build -o $(BIN)/$(API_SERVER_NAME) ./server/cmd/mukgo-api

log: ## Build log server
	go build -o $(BIN)/$(LOG_SERVER_NAME) ./server/cmd/mukgo-log

clean: ## Clean binaries
	rm -f $(BIN)/$(API_SERVER_NAME)
	rm -f $(BIN)/$(LOG_SERVER_NAME)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'