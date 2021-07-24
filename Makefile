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

.PHONY: gen-api
gen-api: ## Generate router and request structs from OpenAPI spec.
	oapi-codegen -generate "chi-server" resource/openapi/user.yaml > gen/openapi/server.go
	oapi-codegen -generate "types" resource/openapi/user.yaml > gen/openapi/type.go
	oapi-codegen -generate "spec" resource/openapi/user.yaml > gen/openapi/spec.go

.PHONY: gen-model
gen-model:
	go generate gen/gen_schema.go

.PHONY: run
run:
	go run cmd/server

.PHONY: test
test:
	go install github.com/kyoh86/richgo
	APP_ENV=test richgo test -cover github.com/gami/layered_arch_example/${TARGET}

.PHONY: lint
lint:
	golangci-lint run