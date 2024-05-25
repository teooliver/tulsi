package bootstrap

import (
	"context"

	"github.com/teooliver/kanban/internal/bootstrap/internal/deps"
	"github.com/teooliver/kanban/internal/config"
)

type AllDeps struct {
	Repos    *deps.Repos
	Services *deps.Services
	Handlers *deps.RestHandlers
}

func Deps(
	ctx context.Context,
	cfg *config.Config,
) (*AllDeps, error) {
	infra, err := deps.InitInfra(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repos := deps.InitRepos(cfg, infra)
	services := deps.InitServices(cfg, infra, repos)
	handlers := deps.InitRestHandlers(cfg, services)

	return &AllDeps{
		Repos:    repos,
		Services: services,
		Handlers: handlers,
	}, nil
}
