package main

import (
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
)

func main() {
	cfg := config.LoadTestConfig()
	databaseConnection := dbconnection.New(&cfg.DatabaseConfig)
	databaseConnection.Connect()
	defer databaseConnection.Close()

	databaseConnection.DeleteMigrations()
}
