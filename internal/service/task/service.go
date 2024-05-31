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
	CreateTask(ctx context.Context, task task.TaskForCreate) error
}

func New(
	task taskRepo,
) *Service {
	return &Service{
		taskRepo: task,
	}
}

func (s *Service) ListAllTasks(ctx context.Context) ([]task.Task, error) {
	tasks, err := s.taskRepo.ListAllTasks(context.TODO())
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Service) CreateTask(ctx context.Context, task task.TaskForCreate) error {
	err := s.taskRepo.CreateTask(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}
