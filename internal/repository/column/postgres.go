package column

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

func (r *PostgresRepository) CreateColumn(ctx context.Context, column ColumnForCreate) (string, error) {
	insertSQL, args, err := goqu.Insert("project_column").Rows(goqu.Record{
		"name":       column.Name,
		"project_id": column.ProjectID,
		"position":   column.Position,
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

func (r *PostgresRepository) UpdateColumn(ctx context.Context, columnID string, column ColumnForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("project_column").Set(column).Where(goqu.Ex{"id": columnID}).Returning("id").ToSQL()
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

func (r *PostgresRepository) GetColumnsByProjectID(ctx context.Context, projectID string) (columns []Column, err error) {
	// SELECT project_column.id, project_column.name,project_column.project_id, project_column.position
	// FROM project_column
	// INNER JOIN project ON project_column.project_id= project.id;
	sql, args, err := goqu.Select(
		"project_column.id",
		"project_column.name",
		"project_column.project_id",
		"project_column.position").From(
		"project_column").InnerJoin(
		goqu.T("project"),
		goqu.On(goqu.Ex{"project_column.project_id": goqu.I("project.id")})).ToSQL()

	if err != nil {
		return []Column{}, fmt.Errorf("error generating update project query: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)

	if err != nil {
		return []Column{}, fmt.Errorf("error executing update project query: %w", err)
	}

	var result []Column
	for rows.Next() {
		row, err := mapRowToColumn(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}

// func (r *PostgresRepository) DeleteColumn(ctx context.Context, columnID string) (string, error) {
// 	insertSQL, args, err := goqu.Delete("project_column").Where(goqu.Ex{"id": columnID}).Returning("id").ToSQL()
// 	if err != nil {
// 		return "", fmt.Errorf("error generating delete column query: %w", err)
// 	}

// 	var id string
// 	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
// 	if err != nil {
// 		return "", fmt.Errorf("error executing delete column query: %w", err)
// 	}

// 	return id, nil
// }
