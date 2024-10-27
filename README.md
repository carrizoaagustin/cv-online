# CV ONLINE
cv-online is a personal project developed in collaboration with Miguel Cabeza, aimed at exposing endpoints to store information related to a person's curriculum vitae, allowing the data to be retrieved for display on a personal presentation page.

## RUN SERVER

### INIT SERVER WITHOUT DOCKER
`go run ./cmd/server`
OR
`make init`

### ADD DEPENDENCIES
`go get package/path`

for example: `go get github.com/gin-gonic/gin`

## HOW TO DEV

### INSTALL DEV TOOLS
`make install-tools`

### MIGRATIONS

#### CREATE MIGRATIONS
`make migrate-create name=<name>`

#### UP MIGRATIONS
`make migrate-up`

#### DOWN MIGRATIONS
`make migrate-down`

obs: delete the last migration applied

### TESTING

#### RUN TESTS
`make test`

#### RENDER HTML COVERAGE
`make cover`

OBS: This command generates an HTML file, which allows you to see which lines of code have been tested so far. You need to have previously run the command `make test`.

## DEPENDENCIES
This project was developed with:
- Gin gonic
- goqu
- google/uuid
- goose
