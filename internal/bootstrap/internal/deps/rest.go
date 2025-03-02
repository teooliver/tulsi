package deps

import (
	"github.com/teooliver/tulsi/internal/config"
	"github.com/teooliver/tulsi/internal/controller/column"
	"github.com/teooliver/tulsi/internal/controller/project"
	"github.com/teooliver/tulsi/internal/controller/status"
	"github.com/teooliver/tulsi/internal/controller/task"
	"github.com/teooliver/tulsi/internal/controller/user"
)

type RestHandlers struct {
	TaskHandler    task.Handler
	UserHandler    user.Handler
	StatusHandler  status.Handler
	ProjectHandler project.Handler
	ColumnHandler  column.Handler
}

func InitRestHandlers(cfg *config.Config, services *Services) *RestHandlers {
	taskHandler := task.New(services.TaskService)
	userHandler := user.New(services.UserService)
	statusHandler := status.New(services.StatusService)
	projectHandler := project.New(services.ProjectService)
	columnHandler := column.New(services.ColumnService)

	return &RestHandlers{
		TaskHandler:    taskHandler,
		UserHandler:    userHandler,
		StatusHandler:  statusHandler,
		ProjectHandler: projectHandler,
		ColumnHandler:  columnHandler,
	}
}
