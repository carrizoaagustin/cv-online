package dbconnection

import (
	"database/sql"
	"embed"
	"log"

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
	dbFullURL := dbs.dbConfig.URL + "/" + dbs.dbConfig.SchemaName + "?sslmode=" + dbs.dbConfig.SSLMode

	databaseConnection, err := sql.Open("postgres", dbFullURL)
	if err != nil {
		panic(err)
	}

	dbs.databaseConnection = databaseConnection
}

func (dbs *DatabaseConnection) CreateSchema() {
	db, err := sql.Open("postgres", dbs.dbConfig.URL+"?sslmode="+dbs.dbConfig.SSLMode)

	if err != nil {
		log.Fatalf("Error DB connection: %v", err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE DATABASE" + dbs.dbConfig.SchemaName)
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}
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
