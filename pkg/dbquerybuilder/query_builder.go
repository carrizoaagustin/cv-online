package dbquerybuilder

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type DBQueryBuilder struct {
	query *goqu.Database
}

func New(dbConnection *sql.DB) *DBQueryBuilder {
	dialect := goqu.Dialect("postgresql")
	db := dialect.DB(dbConnection)

	return &DBQueryBuilder{
		query: db,
	}
}

func (q *DBQueryBuilder) StartQuery() *goqu.Database {
	return q.query
}
