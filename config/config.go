package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	URL        string `env:"PSQL_URL"`
	SchemaName string `env:"PSQL_SCHEMA"`
	SSLMode    string `env:"PSQL_SSL_MODE" envDefault:"disable"`
}

type App struct {
	EnvironmentMode string `env:"APP_ENVIRONMENT" envDefault:"DEV"` // DEV OR PROD
	PORT            string `env:"APP_PORT" envDefault:"8000"`
}

type config struct {
	DatabaseConfig
	App
}

const (
	DevelopmentMode = "DEV"
	ProductionMode  = "PROD"
)

// revive:disable:unexported-return
func LoadConfig() *config {
	if os.Getenv("APP_ENVIRONMENT") != ProductionMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file with godotenv: %v", err)
		}
	}

	var cfg config
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Error parsing env vars: %v", err)
	}

	return &cfg
}

func LoadTestConfig() *config {
	if os.Getenv("APP_ENVIRONMENT") != ProductionMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file with godotenv: %v", err)
		}
	}

	schemaTest := "test-" + os.Getenv("PSQL_SCHEMA")

	os.Setenv("PSQL_SCHEMA", schemaTest)

	var cfg config
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Error parsing env vars: %v", err)
	}

	return &cfg
}
