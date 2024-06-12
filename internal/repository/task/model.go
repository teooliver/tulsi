package task

import (
	"database/sql"
	"net/http"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

type TaskForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusID    string `json:"status_id"`
	Color       string `json:"color"`
	UserID      string `json:"user_id"`
}

type TaskForUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusID    string `json:"status_id"`
	Color       string `json:"color"`
	UserID      string `json:"user_id"`
}

// impl chi.Render interface to render objects as JSON to the API consumer
func (t *Task) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.StatusID, &t.UserID)

	if err != nil {
		// TODO: handle error
		// if err.Is(err, sql.ErrNoRows) {
		// 	return Task{}, serrors.ErrNotFound
		// }
		// return Task{}, errors.Wrap(err, "failed to scan metadata")
	}
	return t, nil
}
