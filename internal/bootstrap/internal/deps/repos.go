package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/repository/column"
	"github.com/teooliver/kanban/internal/repository/project"
	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/internal/repository/user"
)

type Repos struct {
	TaskRepo     *task.PostgresRepository
	StatusRepo   *status.PostgresRepository
	UserRepo     *user.PostgresRepository
	ProjectsRepo *project.PostgresRepository
	ColumnRepo   *column.PostgresRepository
}

func InitRepos(cfg *config.Config, infra *Infra) *Repos {
	taskRepo := task.NewPostgres(infra.Postgres)
	statusRepo := status.NewPostgres(infra.Postgres)
	userRepo := user.NewPostgres(infra.Postgres)
	projectsRepo := project.NewPostgres(infra.Postgres)
	columnRepo := column.NewPostgres(infra.Postgres)

	return &Repos{
		TaskRepo:     taskRepo,
		StatusRepo:   statusRepo,
		UserRepo:     userRepo,
		ProjectsRepo: projectsRepo,
		ColumnRepo:   columnRepo,
	}
}
