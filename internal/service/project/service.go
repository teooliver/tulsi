package project

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/project"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type Service struct {
	projectRepo projectRepo
}

type projectRepo interface {
	ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[project.Project], error)
	// CreateProject(ctx context.Context, project project.ProjectToCreate) (err error)
	// UpdateProject(ctx context.Context, projectID string, project project.ProjectToUpdate) (err error)
	// DeleteProject(ctx context.Context, projectID string) (err error)
}

func New(
	project projectRepo,
) *Service {
	return &Service{
		projectRepo: project,
	}
}

func (s *Service) ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[project.Project], error) {
	allProjects, err := s.projectRepo.ListAllProjects(ctx, params)
	if err != nil {
		return postgresutils.Page[project.Project]{}, err
	}

	return allProjects, nil
}

// func (s *Service) CreateProject(ctx context.Context, project project.ProjectToCreate) error {
// 	err := s.projectRepo.CreateProject(ctx, project)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) DeleteProject(ctx context.Context, statusID string) error {
// 	err := s.projectRepo.DeleteProject(ctx, statusID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (s *Service) UpdateProject(ctx context.Context, statusID string, updatedProject project.ProjectToUpdate) error {
// 	err := s.projectRepo.UpdateProject(ctx, statusID, updatedProject)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
