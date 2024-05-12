package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/teooliver/kanban/internal/config"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Config(ctx context.Context, filenames ...string) (*config.Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		// should ignore error since vars can be set in k8s
		log.Fatal("Error loading .env file")
		// log.Error(ctx, "failed to load env files", err)
	}

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("process config: %w", err)
	}

	return &cfg, nil
}
