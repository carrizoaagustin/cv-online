package dbmigrate

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

func Up(db *sql.DB, dbDialect string, embedMigrations embed.FS, dir string) {
	goose.SetBaseFS(embedMigrations)
	

	if err := goose.SetDialect(dbDialect); err != nil {
			panic(err)
	}

	if err := goose.Up(db, dir); err != nil {
			panic(err)
	}
}

func Down(db *sql.DB, dbDialect string, embedMigrations embed.FS, dir string) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dbDialect); err != nil {
			panic(err)
	}

	if err := goose.Down(db, dir); err != nil {
			panic(err)
	}
}