package task

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllTasks(ctx context.Context) ([]Task, error) {
	sql, _, _ := goqu.From("task").ToSql()

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []Task
	for rows.Next() {
		task, err := mapRowToTask(rows)
		if err != nil {
			// TODO: Handle error
			fmt.Println(err)
		}
		result = append(result, task)
	}

	return result, nil
}
