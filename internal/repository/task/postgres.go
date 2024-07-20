package task

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

func (r *PostgresRepository) ListAllTasks(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[Task], error) {
	q := goqu.From("task").Select(allColumns...)
	return postgresutils.ListPaginated(ctx, r.db, q, params, mapRowToTask)
}

func (r *PostgresRepository) CreateTask(ctx context.Context, task TaskForCreate) (string, error) {
	insertSQL, args, err := goqu.Insert("task").Rows(TaskForCreate{
		Title:       task.Title,
		Description: task.Description,
		Color:       task.Color,
		// UserID:      task.UserID,
	}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating create task query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create task query: %w", err)
	}

	return id, nil
}

// TODO: use `uuid` type for taskID instead of `string`
func (r *PostgresRepository) DeleteTask(ctx context.Context, taskID string) (string, error) {
	insertSQL, args, err := goqu.Delete("task").Where(goqu.Ex{"id": taskID}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating delete task query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create task query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) UpdateTask(ctx context.Context, taskID string, task TaskForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("task").Set(task).Where(goqu.Ex{"id": taskID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update task query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update task query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}

func (r *PostgresRepository) InsertMultipleTasks(ctx context.Context, tasks []TaskForCreate) (err error) {
	insertSQL, args, err := goqu.Insert("task").Rows(tasks).ToSQL()
	if err != nil {
		return fmt.Errorf("error generating insert multiple tasks query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing insert multiple tasks query: %w", err)
	}

	return nil
}
