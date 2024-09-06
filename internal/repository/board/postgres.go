package board

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

func (r *PostgresRepository) ListAllBoards(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[Board], error) {
	q := goqu.From("board").Select(allColumns...)
	return postgresutils.ListPaginated(ctx, r.db, q, params, mapRowToBoard)
}

func (r *PostgresRepository) CreateBoard(ctx context.Context, board BoardForCreate) (string, error) {
	insertSQL, args, err := goqu.Insert("board").Rows(BoardForCreate{
		Name: board.Name,
	}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating create board query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing create board query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) DeleteBoard(ctx context.Context, boardID string) (string, error) {
	insertSQL, args, err := goqu.Delete("board").Where(goqu.Ex{"id": boardID}).Returning("id").ToSQL()
	if err != nil {
		return "", fmt.Errorf("error generating delete board query: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, insertSQL, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error executing delete board query: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) UpdateBoard(ctx context.Context, boardID string, board BoardForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("board").Set(board).Where(goqu.Ex{"id": boardID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update board query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update board query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}
