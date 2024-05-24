package task

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/task"
)

type taskRepo interface {
	ListAllTasks(ctx context.Context) ([]task.Task, error)
	// Get(ctx context.Context, id string) (task.Metadata, error)
	// Create(ctx context.Context, params task.CreateParams)
}

func ListAllTasks(ctx context.Context) ([]task.Task, error) {

}
