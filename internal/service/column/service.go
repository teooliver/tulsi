package project

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/column"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type Service struct {
	columnRepo columnRepo
}

type columnRepo interface {
	ListAllColumns(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[column.Column], error)
	CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error)
}

func New(
	columnRepo columnRepo,
) *Service {
	return &Service{
		columnRepo: columnRepo,
	}
}

func (s *Service) ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[column.Column], error) {
	allColumns, err := s.columnRepo.CreateColumn(ctx, params)
	if err != nil {
		return postgresutils.Page[column.Column]{}, err
	}

	return allColumns, nil
}

func (s *Service) CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error) {
	id, err := s.columnRepo.CreateColumn(ctx, column)
	if err != nil {
		return "", err
	}

	return id, nil
}
