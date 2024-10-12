package config

import "github.com/caarlos0/env/v11"

type DatabaseConfig struct {
	URL string `env:"PSQL_URL"`
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
	var cfg config
	err := env.Parse(&cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}
