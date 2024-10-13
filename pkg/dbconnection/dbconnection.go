package dbconnection

import (
	"database/sql"
	"embed"

	_ "github.com/lib/pq" // PSQL DRIVER.

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbmigrate"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

type DatabaseConnection struct {
	dbConfig           *config.DatabaseConfig
	databaseConnection *sql.DB
}

func New(dbConfig *config.DatabaseConfig) *DatabaseConnection {
	return &DatabaseConnection{
		dbConfig: dbConfig,
	}
}

func (dbs *DatabaseConnection) Connect() {
	databaseConnection, err := sql.Open("postgres", dbs.dbConfig.URL)
	if err != nil {
		panic(err)
	}

	dbs.databaseConnection = databaseConnection
}

func (dbs *DatabaseConnection) Close() {
	dbs.databaseConnection.Close()
}

func (dbs *DatabaseConnection) RunMigrations() {
	dbmigrate.Up(dbs.databaseConnection, "postgres", migrationFiles, "migrations")
}

func (dbs *DatabaseConnection) GetDatabaseConnection() *sql.DB {
	return dbs.databaseConnection
}
