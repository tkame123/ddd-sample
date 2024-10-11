package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	OrderAPIDSN string `env:"ORDER_API_DSN"`
}

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
