package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/service/task"
)

type Services struct {
	TaskService *task.Service
}

func InitServices(cfg *config.Config, infra *Infra, repos *Repos) *Services {
	taskService := task.New(repos.TaskRepo)

	return &Services{
		TaskService: taskService,
	}
}
