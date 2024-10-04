package config

import "github.com/caarlos0/env/v11"


type DatabaseConfig struct {
	Url string `env:"PSQL_URL"`
}


type config struct {
	DatabaseConfig
}

func LoadConfig() *config {
	var cfg config
	env.Parse(&cfg)
	
	return &cfg
}