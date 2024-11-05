package task

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *PostgresRepository) GetTaskByID(ctx context.Context, taskID string) (Task, error) {
	q := goqu.From("task").Select(allColumns...).Where(goqu.Ex{"id": taskID})
	query, args, err := q.ToSQL()
	if err != nil {
		println("TO SQL ERROR")
		return Task{}, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		println("QueryContext ERROR")
		return Task{}, err
	}

	var t Task
	// From the docs:
	// https://pkg.go.dev/database/sql#Rows.Next
	// Every call to Rows.Scan, even the first one, must be preceded by a call to Rows.Next.
	row.Next()
	err = row.Scan(&t.ID, &t.Title, &t.Description, &t.Color, &t.StatusID, &t.UserID)

	if err != nil {
		return Task{}, fmt.Errorf("Error error scanning Task row: %w", err)
	}

	return t, nil
}

func (r *PostgresRepository) CreateTask(ctx context.Context, task TaskForCreate) (string, error) {
	insertSQL, args, err := goqu.Insert("task").Rows(goqu.Record{
		"title":       task.Title,
		"description": task.Description,
		"color":       task.Color,
		"column_id":   task.ColumnID,
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
		return "", fmt.Errorf("error executing delete task query: %w", err)
	}

	return id, nil
}

// TODO: Return the result from ExecContext
func (r *PostgresRepository) UpdateTask(ctx context.Context, taskID string, task TaskForUpdate) (err error) {
	// TODO: update only the fields sent by the FE instead of using the whole task blindly:
	// SET column1 = value1, column2 = value2, ...
	updateSQL, args, err := goqu.Update("task").Set(task).Where(goqu.Ex{"id": taskID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update task query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update task query: %w", err)
	}

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
