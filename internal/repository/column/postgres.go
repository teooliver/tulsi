package column

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

func (r *PostgresRepository) ListAllColumns(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[Column], error) {
	q := goqu.From("column").Select(allColumns...)
	return postgresutils.ListPaginated(ctx, r.db, q, params, mapRowToColumn)
}

func (r *PostgresRepository) CreateColumn(ctx context.Context, column ColumnForCreate) (string, error) {
	insertSQL, args, err := goqu.Insert("column").Rows(ColumnForCreate{
		Name:      column.Name,
		ProjectID: column.ProjectID,
		Position:  column.Position,
	}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating create column query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create column query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) DeleteColumn(ctx context.Context, columnID string) (string, error) {
	insertSQL, args, err := goqu.Delete("column").Where(goqu.Ex{"id": columnID}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating delete column query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing delete column query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) UpdateColumn(ctx context.Context, columnID string, column ColumnForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("column").Set(column).Where(goqu.Ex{"id": columnID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update column query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update column query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}
