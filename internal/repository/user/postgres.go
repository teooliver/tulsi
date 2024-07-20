package user

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllUsers(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[User], error) {
	q := goqu.From("app_user").Select(allColumns...)
	return postgresutils.ListPaginated(ctx, r.db, q, params, mapRowToUser)
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user UserForCreate) (err error) {
	insertSQL, args, err := goqu.Insert("app_user").Rows(UserForCreate{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating create user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing create user query: %w", err)
	}

	slog.Info("CREATED RESULT", result)
	return nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, userID string) (err error) {
	insertSQL, args, err := goqu.Delete("app_user").Where(goqu.Ex{"id": userID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating delete user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing delete user query: %w", err)
	}

	slog.Info("DELETED USER ID", result)
	return nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, userID string, user UserForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("app_user").Set(user).Where(goqu.Ex{"id": userID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update user query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}
