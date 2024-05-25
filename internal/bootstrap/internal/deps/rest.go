package deps

import (
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/internal/controller/task"
)

type RestHandlers struct {
	TaskHandler task.Handler
}

func InitRestHandlers(cfg *config.Config, services *Services) *RestHandlers {

	taskHandler := task.New(services.TaskService)

	return &RestHandlers{
		TaskHandler: taskHandler,
	}
}
