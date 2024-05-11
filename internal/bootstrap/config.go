package bootstrap

import (
	"context"
	"fmt"

	"github.com/teooliver/kanban/internal/config"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Config(ctx context.Context, filenames ...string) (*config.Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		// log.Error(ctx, "failed to load env files", err)
		// ignore error since vars can be set in k8s
	}

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("process config: %w", err)
	}

	return &cfg, nil
}
