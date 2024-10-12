package dbconnection

import (
	"database/sql"
	"embed"

	// Driver psql.
	_ "github.com/lib/pq"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbmigrate"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func ConnectDB(dbConfig *config.DatabaseConfig) *sql.DB {
	db, err := sql.Open("postgres", dbConfig.URL)
	if err != nil {
		panic(err)
	}

	return db
}

func RunMigrations(db *sql.DB) {
	dbmigrate.Up(db, "postgres", migrationFiles, "migrations")
}
