// Package config loads server configuration from env; invalid config aborts startup (Fail Fast).
package config

import (
	"fmt"
	"os"
	"strconv"
)

type Env string

const (
	EnvDev  Env = "dev"
	EnvProd Env = "prod"
)

type Config struct {
	Port        int
	Env         Env
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config{Port: 8080, Env: EnvDev}

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("config: DATABASE_URL is required")
	}

	if v := os.Getenv("PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil || p < 1 || p > 65535 {
			return Config{}, fmt.Errorf("config: invalid PORT %q", v)
		}
		cfg.Port = p
	}

	if v := os.Getenv("APP_ENV"); v != "" {
		switch Env(v) {
		case EnvDev, EnvProd:
			cfg.Env = Env(v)
		default:
			return Config{}, fmt.Errorf("config: invalid APP_ENV %q (want dev|prod)", v)
		}
	}

	return cfg, nil
}
