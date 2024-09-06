package board

import (
	"database/sql"
	"fmt"
)

type Board struct {
	ID   string `json:"id"` //PK
	Name string `json:"name"`
	// IsPublic bool   `json:"is_public"`
	// UserID string   `json:user_id` //FK
	// CreatedDate string `json:"created_date"`
}
type BoardForCreate struct {
	Name string `json:"name"`
}

type BoardForUpdate struct {
	Name string `json:"name"`
}

var allColumns = []any{
	"id",
	"name",
}

func mapRowToBoard(rows *sql.Rows) (Board, error) {
	var t Board
	err := rows.Scan(&t.ID, &t.Name)

	if err != nil {
		return Board{}, fmt.Errorf("Error error scanning Task row: %w", err)

	}
	return t, nil
}
