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

	tasks, err := s.taskRepo.ListAllTasks(context.TODO())
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
