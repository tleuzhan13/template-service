package config

import (
	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload" // lib for autoload env files
	"template-service/pkg/mongo"
)

type (
	Config struct {
		Log   Log
		Mongo mongo.Config
		GRPC  GRPC
	}

	Log struct {
		Level   string `env:"LOG_LEVEL,required"`
		Mode    string `env:"LOG_MODE,required"`
		License string `env:"LICENSE,required"`
	}

	GRPC struct {
		Port           int  `env:"GRPC_PORT,required"`
		RequestLogging bool `env:"GRPC_REQUEST_LOGGING"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
