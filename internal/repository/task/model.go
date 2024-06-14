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

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Color, &t.Status, &t.StatusID, &t.UserID)

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
