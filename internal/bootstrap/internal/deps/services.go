package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/service/status"
	"github.com/teooliver/kanban/internal/service/task"
	"github.com/teooliver/kanban/internal/service/user"
)

type Services struct {
	TaskService   *task.Service
	UserService   *user.Service
	StatusService *status.Service
}

func InitServices(cfg *config.Config, infra *Infra, repos *Repos) *Services {
	taskService := task.New(repos.TaskRepo)
	userService := user.New(repos.UserRepo)
	statusService := status.New(repos.StatusRepo)

	return &Services{
		TaskService:   taskService,
		UserService:   userService,
		StatusService: statusService,
	}
}
