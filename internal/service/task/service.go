package task

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type Service struct {
	taskRepo taskRepo
}

type taskRepo interface {
	ListAllTasks(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[task.Task], error)
	GetTaskByID(ctx context.Context, taskID string) (task.Task, error)
	CreateTask(ctx context.Context, task task.TaskForCreate) (string, error)
	DeleteTask(ctx context.Context, taskID string) (string, error)
	UpdateTask(ctx context.Context, taskID string, task task.TaskForUpdate) (err error)
}

func New(
	task taskRepo,
) *Service {
	return &Service{
		taskRepo: task,
	}
}

func (s *Service) ListAllTasks(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[task.Task], error) {
	tasks, err := s.taskRepo.ListAllTasks(ctx, params)
	if err != nil {
		return postgresutils.Page[task.Task]{}, err
	}

	return tasks, nil
}

func (s *Service) GetTaskByID(ctx context.Context, taskID string) (task.Task, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, taskID)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *Service) CreateTask(ctx context.Context, task task.TaskForCreate) (string, error) {
	id, err := s.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) UpdateTask(ctx context.Context, taskID string, updatedTask task.TaskForUpdate) error {
	err := s.taskRepo.UpdateTask(ctx, taskID, updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, taskID string) (string, error) {
	id, err := s.taskRepo.DeleteTask(ctx, taskID)
	if err != nil {
		return "", err
	}

	return id, nil
}
