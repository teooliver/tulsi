package task

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListTasks() {
	sql, _, _ := goqu.From("task").ToSql()

	rows, err := r.db.Query(sql)
	if err != nil {
		log.Fatal("Error listing tasks")
	}

	defer rows.Close()

	for rows.Next() {
		task, _ := mapRowToTask(rows)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// CheckError(err)

		fmt.Println(task)
	}

	// return tasks
}
