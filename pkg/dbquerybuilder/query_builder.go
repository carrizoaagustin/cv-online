package dbquerybuilder

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"

	"github.com/carrizoaagustin/cv-online/pkg/casefmt"
)

type DBQueryBuilder struct {
	query *goqu.Database
}

func New(dbConnection *sql.DB) *DBQueryBuilder {
	dialect := goqu.Dialect("postgresql")
	db := dialect.DB(dbConnection)
	goqu.SetColumnRenameFunction(casefmt.CamelCaseToSnakeCase)

	return &DBQueryBuilder{
		query: db,
	}
}

func (q *DBQueryBuilder) StartQuery() *goqu.Database {
	return q.query
}

type Record map[string]interface{}
