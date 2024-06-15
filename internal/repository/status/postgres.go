package status

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllStatus(ctx context.Context) ([]Status, error) {
	sql, _, err := goqu.From("status").Select(allColumns...).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("error generating list all status query: %w", err)
	}

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("error executing list all status query: %w", err)
	}

	defer rows.Close()

	var result []Status
	for rows.Next() {
		task, err := mapRowToStatus(rows)
		if err != nil {
			return nil, err
		}
		slog.Info("LIST SQL RESULT ===> %+v\n", "result", task)

		result = append(result, task)
	}

	return result, nil
}

func (r *PostgresRepository) CreateStatus(ctx context.Context, status StatusForCreate) (err error) {
	insertSQL, args, err := goqu.Insert("task").Rows(StatusForCreate{
		Name: status.Name,
	}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating create status query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing create status query: %w", err)
	}

	slog.Info("CREATE RESULT", result)
	return nil
}

func (r *PostgresRepository) DeleteStatus(ctx context.Context, statusID string) (err error) {
	insertSQL, args, err := goqu.Delete("task").Where(goqu.Ex{"id": statusID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating delete status query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing delete status query: %w", err)
	}

	slog.Info("DELETED TASKS ID", result)
	return nil
}

func (r *PostgresRepository) UpdateStatus(ctx context.Context, statusID string, status StatusForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("task").Set(status).Where(goqu.Ex{"id": statusID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update status query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update status query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}
