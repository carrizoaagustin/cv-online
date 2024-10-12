# DEVELOPMENT

## BASIC COMMANDS

### INIT SERVER WITHOUT DOCKER
`go run ./cmd/server`

### ADD DEPENDENCIES
`go get package/path`

for example: `go get github.com/gin-gonic/gin`


## DEPENDENCIES
- Gin gonic
- goqu
- google/uuid
- goose (migrations) (install)


## GOOSE
https://github.com/pressly/goose

### INSTALL
`go install github.com/pressly/goose/v3/cmd/goose@latest`

### CREATE MIGRATION
`GOOSE_MIGRATION_DIR=./pkg/dbconnection/migrations goose create <name> sql`

### DOWN

`GOOSE_MIGRATION_DIR=./pkg/dbconnection/migrations GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:root1234@localhost:5432/example?sslmode=disable" goose down`

### UP

`GOOSE_MIGRATION_DIR=./pkg/dbconnection/migrations GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:root1234@localhost:5432/example?sslmode=disable" goose up`