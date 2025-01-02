package dbconnection

import (
	"database/sql"
	"embed"
	"fmt"
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
	url := fmt.Sprintf("%s?sslmode=%s", dbs.dbConfig.URL, dbs.dbConfig.SSLMode)
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("Error DB connection: %v", err)
	}

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1);`
	err = db.QueryRow(query, dbs.dbConfig.SchemaName).Scan(&exists)

	if err != nil {
		log.Fatalf("Failed to check database existence: %v", err)
	}

	queryCreateSchema := fmt.Sprintf(`CREATE DATABASE "%s"`, dbs.dbConfig.SchemaName)

	if !exists {
		_, err = db.Exec(queryCreateSchema)
		if err != nil {
			log.Fatalf("Error creating schema: %v", err)
		}
	}

	defer db.Close()
}

func (dbs *DatabaseConnection) Close() {
	dbs.databaseConnection.Close()
}

func (dbs *DatabaseConnection) RunMigrations() {
	dbmigrate.Up(dbs.databaseConnection, "postgres", migrationFiles, "migrations")
}

func (dbs *DatabaseConnection) DeleteMigrations() {
	dbmigrate.Down(dbs.databaseConnection, "postgres", migrationFiles, "migrations")
}

func (dbs *DatabaseConnection) GetDatabaseConnection() *sql.DB {
	return dbs.databaseConnection
}
