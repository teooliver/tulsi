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

func New(
	task taskRepo,
) *Service {
	return &Service{
		taskRepo: task,
	}
}

func (s *Service) ListAllTasks(ctx context.Context) ([]task.Task, error) {
	// grab data from repo
	//
	test_task := task.Task{
		Title: "hello there title",
	}
	tasks := make([]task.Task, 0)
	tasks = append(tasks, test_task)

	return tasks, nil
}
