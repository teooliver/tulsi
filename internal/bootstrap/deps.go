package bootstrap

import (
	"context"

	"github.com/teooliver/tulsi/internal/bootstrap/internal/deps"
	"github.com/teooliver/tulsi/internal/config"
)

type AllDeps struct {
	// logger
	// 	config? config
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
		// Do we need to expose the repos here?
		Repos: repos,
		// Do we need to expose the services here?
		Services: services,
		Handlers: handlers,
	}, nil
}
