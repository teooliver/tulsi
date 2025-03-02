package project

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/teooliver/tulsi/pkg/postgresutils"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[Project], error) {
	q := goqu.From("project").Select(allColumns...)
	return postgresutils.ListPaginated(ctx, r.db, q, params, mapRowToProject)
}

func (r *PostgresRepository) CreateProject(ctx context.Context, project CreateProjectRequest) (string, error) {

	insertSQL, args, err := goqu.Insert("project").Rows(goqu.Record{
		"name":        project.Name,
		"description": project.Description,
		"is_archived": false,
	}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating create project query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create project query: %w", err)
	}

	return id, nil
}

// TODO: use `uuid` type for projectID instead of `string`
func (r *PostgresRepository) ArquiveProject(ctx context.Context, projectID string) (string, error) {
	updateSQL, args, err := goqu.Update("project").Set(
		goqu.Record{"is_archived": true}).Where(
		goqu.Ex{"id": projectID}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error arquiving project: %w", err)
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return "", fmt.Errorf("error executing arquive project query: %w", err)
	}

	// TODO: return id here
	return "", nil
}

// TODO: Return the result from ExecContext
func (r *PostgresRepository) UpdateProject(ctx context.Context, projectID string, project ProjectToUpdate) (err error) {
	updateSQL, args, err := goqu.Update("project").Set(project).Where(goqu.Ex{"id": projectID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update project query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update project query: %w", err)
	}

	// slog.Info("UPDATED ID", result)
	return nil
}
