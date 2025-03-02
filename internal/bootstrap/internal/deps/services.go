package deps

import (
	"github.com/teooliver/tulsi/internal/config"
	"github.com/teooliver/tulsi/internal/service/column"
	"github.com/teooliver/tulsi/internal/service/project"
	"github.com/teooliver/tulsi/internal/service/status"
	"github.com/teooliver/tulsi/internal/service/task"
	"github.com/teooliver/tulsi/internal/service/user"
)

type Services struct {
	TaskService    *task.Service
	UserService    *user.Service
	StatusService  *status.Service
	ProjectService *project.Service
	ColumnService  *column.Service
}

func InitServices(cfg *config.Config, infra *Infra, repos *Repos) *Services {
	taskService := task.New(repos.TaskRepo)
	userService := user.New(repos.UserRepo)
	statusService := status.New(repos.StatusRepo)
	projectService := project.New(repos.ProjectsRepo)
	columnService := column.New(repos.ColumnRepo)

	return &Services{
		TaskService:    taskService,
		UserService:    userService,
		StatusService:  statusService,
		ProjectService: projectService,
		ColumnService:  columnService,
	}
}
