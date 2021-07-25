PWD:=$(shell pwd)
TARGET=...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Install depeendent tools and setup project
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.1
	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.6.0
	go get github.com/pressly/goose@v2.7.0
	go mod tidy

.PHONY: gen-api
gen-api: ## Generate router and request type structs from OpenAPI spec.
	go generate gen/gen_openapi.go

.PHONY: gen-model
gen-model: ## Genereate SQLBoiler models
	go generate gen/gen_schema.go

.PHONY: run
run: ## Run local API server
	go run cmd/server

.PHONY: test
test: ## Run test API server
	go install github.com/kyoh86/richgo
	APP_ENV=test richgo test -cover app/${TARGET}

.PHONY: lint
lint: ## Run linter
	golangci-lint run


.PHONY: setup_test_db
setup_test_db: ## Setup test db
