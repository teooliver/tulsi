package column

import (
	"database/sql"
	"fmt"
)

type Column struct {
	ID        string `json:"id"` //PK
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  string `json:"position"`
}

type ColumnForCreate struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  string `json:"position"`
}

type ColumnForUpdate struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  string `json:"position"`
}

var allColumns = []any{
	"id",
	"name",
	"project_id",
	"position",
}

func mapRowToColumn(rows *sql.Rows) (Column, error) {
	var t Column
	err := rows.Scan(&t.ID, &t.Name)

	if err != nil {
		return Column{}, fmt.Errorf("Error error scanning Task row: %w", err)

	}
	return t, nil
}
