package main

import (
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
)

func main() {
	cfg :=config.LoadConfig()
	db := dbconnection.ConnectDB(&cfg.DatabaseConfig)
	defer db.Close()

	dbconnection.RunMigrations(db)

	// run app
}