package project

import (
	"context"

	"github.com/teooliver/tulsi/internal/repository/project"
	"github.com/teooliver/tulsi/pkg/postgresutils"
)

type Service struct {
	projectRepo projectRepo
}

type projectRepo interface {
	ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[project.Project], error)
	ArquiveProject(ctx context.Context, projectID string) (string, error)
	CreateProject(ctx context.Context, project project.CreateProjectRequest) (string, error)
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

func (s *Service) CreateProject(ctx context.Context, project project.CreateProjectRequest) (string, error) {
	id, err := s.projectRepo.CreateProject(ctx, project)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) ArquiveProject(ctx context.Context, projectID string) (string, error) {
	id, err := s.projectRepo.ArquiveProject(ctx, projectID)
	if err != nil {
		return id, err
	}
	return id, nil
}

// func (s *Service) UpdateProject(ctx context.Context, statusID string, updatedProject project.ProjectToUpdate) error {
// 	err := s.projectRepo.UpdateProject(ctx, statusID, updatedProject)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
