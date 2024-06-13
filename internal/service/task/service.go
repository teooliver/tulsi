package task

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/pkg/seedDb"
)

type Service struct {
	taskRepo taskRepo
}

type taskRepo interface {
	ListAllTasks(ctx context.Context) ([]task.Task, error)
	CreateTask(ctx context.Context, task task.TaskForCreate) error
	DeleteTask(ctx context.Context, taskID string) error
	InsertMultipleTasks(ctx context.Context, tasks []task.TaskForCreate) error
	UpdateTask(ctx context.Context, taskID string, task task.TaskForUpdate) (err error)
}

func New(
	task taskRepo,
) *Service {
	return &Service{
		taskRepo: task,
	}
}

func (s *Service) ListAllTasks(ctx context.Context) ([]task.Task, error) {
	tasks, err := s.taskRepo.ListAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Service) CreateTask(ctx context.Context, task task.TaskForCreate) error {
	err := s.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, taskID string) error {
	err := s.taskRepo.DeleteTask(ctx, taskID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateTask(ctx context.Context, taskID string, updatedTask task.TaskForUpdate) error {
	err := s.taskRepo.UpdateTask(ctx, taskID, updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) InsertMultipleTasks(ctx context.Context) error {
	newTasks := seedDb.CreateMultipleTasks(20)
	err := s.taskRepo.InsertMultipleTasks(ctx, newTasks)
	if err != nil {
		return err
	}

	return nil
}
