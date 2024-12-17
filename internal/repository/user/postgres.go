package user

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
	"github.com/teooliver/kanban/pkg/auth"
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

// TODO: Return created user instead of just the id
func (r *PostgresRepository) CreateUser(ctx context.Context, user UserForCreate) (string, error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	insertSQL, args, err := goqu.Insert("app_user").Rows(goqu.Record{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"login": goqu.Record{
			"hashedPassword": hashedPassword,
			"csrf_token":     "",
			"session_token":  "",
		},
	}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating create user query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create task query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, userID string) (userId string, err error) {
	insertSQL, args, err := goqu.Delete("app_user").Where(goqu.Ex{"id": userID}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating delete user query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing delete user query: %w", err)
	}

	return id, nil

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

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	q := goqu.From("app_user").Select(allColumns...).Where(goqu.Ex{"email": email})
	query, args, err := q.ToSQL()
	if err != nil {
		return User{}, fmt.Errorf("error generating get user query: %w", err)
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return User{}, fmt.Errorf("error executing get user query: %w", err)
	}

	var u User
	err = row.Scan(u.ID, u.Email, u.FirstName, u.LastName, u.Login.HashedPassword, u.Login.SessionToken, u.Login.CSRFToken)

	if err != nil {
		return User{}, fmt.Errorf("Error error scanning Task row: %w", err)
	}

	return u, nil

}
