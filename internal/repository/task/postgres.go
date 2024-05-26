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
	sql, _, err := goqu.From("task").ToSql()
	if err != nil {
		fmt.Print("TO SQL ERROR")
	}

	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Println(err)
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

func (r *PostgresRepository) CreateTask(ctx context.Context, task Task) (id string, err error) {
	// "INSERT INTO task (
	//            title,
	//            description,
	//            status_id,
	//            color,
	//            user_id
	//            )
	//            values ($1, $2,$3, $4, $5) RETURNING id",
	//
	//
	// ds := goqu.Insert("task").
	// 	Cols("first_name", "last_name").
	// 	Vals(
	// 		goqu.Vals{"Greg", "Farley"},
	// 		goqu.Vals{"Jimmy", "Stewart"},
	// 		goqu.Vals{"Jeff", "Jeffers"},
	// 	)
	// insertSQL, args, _ := ds.ToSQL()

	// sql, _, err := goqu.From("task").ToSql()
}
