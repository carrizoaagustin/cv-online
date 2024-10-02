package config_migrations

import (
	"embed"
)

//go:embed migrations/*.sql
var EmbedMigrations embed.FS