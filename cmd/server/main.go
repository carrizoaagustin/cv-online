package main

import (
	"database/sql"
	"log"

	config_migrations "github.com/carrizoaagustin/cv-online"
	"github.com/carrizoaagustin/cv-online/internal/user/infrastructure/database"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)




func main() {
	database.HelloWorld()

  // Establecer la conexión a la base de datos
	db, err := sql.Open("postgres", "user=postgres password=root1234 dbname=example sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()

	goose.SetBaseFS(config_migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
			panic(err)
	}

	// run app
}