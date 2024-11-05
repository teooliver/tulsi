package task

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	StatusID    *string `json:"status_id"`
	ColumnID    string  `json:"column_id"`
	UserID      *string `json:"user_id"`
	// CreatedDate string  `json:"created_date"`
	// IsActive bool `json:"is_active"`
}

type TaskForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	ColumnID    string `json:"column_id"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
	// ColumnID string  `json:"column_id"`
	// CreatedDate string  `json:"created_date"`
	// IsActive bool `json:"is_active"`
}

type TaskForUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	ColumnID    string `json:"column_id"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
	// CreatedDate string  `json:"created_date"`
}

var allColumns = []any{
	"id",
	"title",
	"description",
	"color",
	"status_id",
	"column_id",
	"user_id",
	// "created_date",
	// "is_active",
}

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Color, &t.StatusID, &t.ColumnID, &t.UserID)

	if err != nil {
		return Task{}, fmt.Errorf("Error error scanning Task row: %w", err)
	}

	return t, nil
}
