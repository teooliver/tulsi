package models

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

type TaskForCreate struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

type TaskForUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}
