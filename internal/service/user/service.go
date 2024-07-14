package user

import (
	"context"

	"github.com/teooliver/kanban/internal/repository/user"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type Service struct {
	userRepo userRepo
}

type userRepo interface {
	ListAllUsers(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[user.User], error)
	CreateUser(ctx context.Context, user user.UserForCreate) error
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, userID string, user user.UserForUpdate) (err error)
}

func New(
	user userRepo,
) *Service {
	return &Service{
		userRepo: user,
	}
}

func (s *Service) ListAllUsers(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[user.User], error) {
	users, err := s.userRepo.ListAllUsers(ctx, params)
	if err != nil {
		return postgresutils.Page[user.User]{}, err
	}

	return users, nil
}

func (s *Service) CreateUser(ctx context.Context, user user.UserForCreate) error {
	err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	err := s.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateUser(ctx context.Context, userID string, updatedUser user.UserForUpdate) error {
	err := s.userRepo.UpdateUser(ctx, userID, updatedUser)
	if err != nil {
		return err
	}

	return nil
}
