package status

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type Service struct {
	statusRepo statusRepo
}

type statusRepo interface {
	ListAllStatus(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[status.Status], error)
	CreateStatus(ctx context.Context, status status.StatusForCreate) error
	DeleteStatus(ctx context.Context, statusId string) error
	UpdateStatus(ctx context.Context, statusID string, status status.StatusForUpdate) (err error)
}

func New(
	status statusRepo,
) *Service {
	return &Service{
		statusRepo: status,
	}
}

func (s *Service) ListAllStatus(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[status.Status], error) {
	allStatus, err := s.statusRepo.ListAllStatus(ctx, params)
	if err != nil {
		return postgresutils.Page[status.Status]{}, err
	}

	return allStatus, nil
}

func (s *Service) CreateStatus(ctx context.Context, status status.StatusForCreate) error {
	err := s.statusRepo.CreateStatus(ctx, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteStatus(ctx context.Context, statusID string) error {
	err := s.statusRepo.DeleteStatus(ctx, statusID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateStatus(ctx context.Context, statusID string, updatedStatus status.StatusForUpdate) error {
	err := s.statusRepo.UpdateStatus(ctx, statusID, updatedStatus)
	if err != nil {
		return err
	}

	return nil
}
