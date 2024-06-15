package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/controller/status"
	"github.com/teooliver/kanban/internal/controller/task"
	"github.com/teooliver/kanban/internal/controller/user"
)

type RestHandlers struct {
	TaskHandler   task.Handler
	UserHandler   user.Handler
	StatusHandler status.Handler
}

func InitRestHandlers(cfg *config.Config, services *Services) *RestHandlers {
	taskHandler := task.New(services.TaskService)
	userHandler := user.New(services.UserService)
	statusHandler := status.New(services.StatusService)

	return &RestHandlers{
		TaskHandler:   taskHandler,
		UserHandler:   userHandler,
		StatusHandler: statusHandler,
	}
}
