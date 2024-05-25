package task

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/task"
)

type Service struct {
	taskRepo taskRepo
}

type taskRepo interface {
	ListAllTasks(ctx context.Context) ([]task.Task, error)
}

func ListAllTasks(ctx context.Context) ([]task.Task, error) {
	return [], nil
}

func New(
	task taskRepo,
) *Service {
	return &Service{
		taskRepo: task,
	}
}
