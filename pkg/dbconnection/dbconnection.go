package dbconnection

import (
	"database/sql"
	"embed"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbmigrate"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func ConnectDB(dbConfig *config.DatabaseConfig) (*sql.DB) {
	db, err := sql.Open("postgres", dbConfig.Url)
	if err != nil {
		panic(err)
	}

	return db
}

func RunMigrations(db *sql.DB) {
    dbmigrate.Up(db, "postgres", migrationFiles, "migrations")
}