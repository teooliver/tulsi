package task

import (
	"database/sql"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	StatusID    *string `json:"status_id"`
	UserID      *string `json:"user_id"`
	Status      *string `json:"status"`
}

type TaskForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
	// Status      *string `json:"status"`
}

type TaskForUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
	// Status      *string `json:"status"`
}

var allColumns = []any{
	"id",
	"title",
	"description",
	"color",
	"status_id",
	"user_id",
	"status",
}

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Color, &t.StatusID, &t.UserID, &t.Status)

	if err != nil {
		return Task{}, err
		// TODO: handle error
		// if err.Is(err, sql.ErrNoRows) {
		// 	return Task{}, serrors.ErrNotFound
		// }
		// return Task{}, errors.Wrap(err, "failed to scan metadata")
	}
	return t, nil
}
