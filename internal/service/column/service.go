package column

import (
	"context"

	"log/slog"

	"github.com/teooliver/tulsi/internal/repository/column"
)

type Service struct {
	columnRepo columnRepo
}

type columnRepo interface {
	CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error)
	GetColumnsByProjectID(ctx context.Context, projectID string) (columns []column.Column, err error)
}

func New(
	columnRepo columnRepo,
) *Service {
	return &Service{
		columnRepo: columnRepo,
	}
}

func (s *Service) CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error) {
	// TODO: check `position` is already taken in the project
	slog.Info("GOT HERE CREATE COLUMN SERVICE")
	id, err := s.columnRepo.CreateColumn(ctx, column)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) GetColumnsByProjectID(ctx context.Context, projectID string) ([]column.Column, error) {
	columns, err := s.columnRepo.GetColumnsByProjectID(ctx, projectID)
	if err != nil {
		return columns, err
	}

	return columns, nil
}
