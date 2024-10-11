package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	OrderAPIDSN string `env:"ORDER_API_DSN"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
