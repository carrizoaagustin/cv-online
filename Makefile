# UPLOAD ENV VARS
include .env
export $(shell sed 's/=.*//' .env)

# CONFIGS
GOBIN ?= $$(go env GOPATH)/bin
MIGRATION_DIR=./pkg/dbconnection/migrations
GOOSE_DRIVER=postgres


.PHONY: install-tools
install-tools:
	go install github.com/vladopajic/go-test-coverage/v2@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./cmd/server

.PHONY: test
test:
	@echo ""
	@echo "---TESTING---"
	@echo ""
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./... -v; 
	@echo ""
	@echo "---COVERAGE OUTPUT---"
	@echo ""
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml
	@echo ""

.PHONY: cover
cover:
	@if [ -f "cover.out" ]; then \
		go tool cover -html=cover.out; \
	else \
		echo "Error: you must first run 'make test'"; \
		exit 1; \
	fi

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then echo "Error: You must provide the migration name with 'name=<name>'"; exit 1; fi
	GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) goose create $(name) sql

.PHONY: migrate-up
migrate-up:
	GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$$PSQL_URL/$$PSQL_SCHEMA?sslmode=$$PSQL_SSL_MODE" goose up

.PHONY: migrate-down
migrate-down:
	GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$$PSQL_URL/$$PSQL_SCHEMA?sslmode=$$PSQL_SSL_MODE" goose down

