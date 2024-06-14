package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/internal/repository/task"
)

type Repos struct {
	TaskRepo   *task.PostgresRepository
	StatusRepo *status.PostgresRepository
}

func InitRepos(cfg *config.Config, infra *Infra) *Repos {
	taskRepo := task.NewPostgres(infra.Postgres)
	statusRepo := status.NewPostgres(infra.Postgres)

	return &Repos{
		TaskRepo:   taskRepo,
		StatusRepo: statusRepo,
	}
}
