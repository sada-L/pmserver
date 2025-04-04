package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App  App
		Http HTTP
		DB   DB
	}

	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" envDefault:"8080"`
	}

	DB struct {
		URL string `env:"DB_URL" envDefault:"postgresql://user:password@localhost:5432/db?sslmode=disable"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
