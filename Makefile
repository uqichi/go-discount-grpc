GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

.DEFAULT_GOAL := help

files: ## List Go files
	@echo $(GOFILES)

packages: ## List Go packages
	@echo $(GOPACKAGES)

test: ## Exec test
	go test -v .

protoc: ## Exec protoc
	protoc --go_out=plugins=grpc:. proto/*.proto

server: ## Run server
	go run server.go

depcmd: ## Install command dependencies
	go get github.com/gobuffalo/pop/soda

up: ## Docker compose up
	docker-compose up -d

down: ## Docker compose down
	docker-compose down

logs: ## Docker compose logs
	@echo not yet

genmodel: ## Generate new model
	@read -p "Model name: " f; \
    soda generate model $${f} --skip-migration

gendbmig: ## Generate new migration
	@read -p "Migration name: " f; \
	soda generate sql $${f}

mig: ## Run database migration
	soda migrate up

migdown: ## Run database migration rollback
	soda migrate down

migr: ## Run database migration reset (down and then up)
	soda migrate reset

release-note: ## Output release note
	@$(SHELL) releasenote.sh

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
