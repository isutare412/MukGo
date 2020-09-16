BIN := bin
API_SERVER_NAME ?= mukgo-api

.DEFAULT_GOAL := help
.PHONY: help api clean

server: api ## Build all server binaries

api: ## Build api server binary
	go build -o $(BIN)/$(API_SERVER_NAME) ./server/cmd/mukgo-api

clean: ## Clean binaries
	rm -f $(BIN)/$(API_SERVER_NAME)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'