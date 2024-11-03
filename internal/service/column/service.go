package column

import (
	"context"

	"log/slog"

	"github.com/teooliver/kanban/internal/repository/column"
)

type Service struct {
	columnRepo columnRepo
}

type columnRepo interface {
	CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error)
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
	slog.Info("GOT HERE %+v\n", column)
	id, err := s.columnRepo.CreateColumn(ctx, column)
	if err != nil {
		return "", err
	}

	return id, nil
}