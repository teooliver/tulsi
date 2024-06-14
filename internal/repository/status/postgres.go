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
		return nil, err
	}

	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var result []Status
	for rows.Next() {
		task, err := mapRowToTask(rows)
		if err != nil {
			// TODO: Handle error
			return nil, err
		}
		slog.Info("LIST SQL RESULT ===> %+v\n", "result", task)

		result = append(result, task)
	}

	return result, nil
}

func (r *PostgresRepository) CreateStatus(ctx context.Context, status StatusForCreate) (err error) {
	insertSQL, args, _ := goqu.Insert("task").Rows(StatusForCreate{
		Name: status.Name,
	}).Returning("id").ToSQL()

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	// TODO: handle error
	if err != nil {
		return err
	}

	slog.Info("CREATE RESULT", result)
	return nil
}

func (r *PostgresRepository) DeleteStatus(ctx context.Context, statusID string) (err error) {
	insertSQL, args, _ := goqu.Delete("task").Where(goqu.Ex{"id": statusID}).Returning("id").ToSQL()

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	// TODO: handle error
	if err != nil {
		fmt.Println(err)
		return err
	}

	slog.Info("DELETED TASKS ID", result)
	return nil
}

func (r *PostgresRepository) UpdateStatus(ctx context.Context, statusID string, status StatusForUpdate) (err error) {
	updateSQL, args, _ := goqu.Update("task").Set(status).Where(goqu.Ex{"id": statusID}).Returning("id").ToSQL()

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	// TODO: handle error
	if err != nil {
		fmt.Println(err)
		return err
	}

	slog.Info("UPDATED ID", result)
	return nil
}
