package column

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type Column struct {
	ID        string `json:"id"` //PK
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  int16  `json:"position"`
}

type ColumnForCreate struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  int16  `json:"position"`
}

type ColumnForUpdate struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` //FK
	Position  int16  `json:"position"`
}

var allColumns = []any{
	"id",
	"name",
	"project_id",
	"position",
}

// TODO: Need this on the Projects repo to map Join result, but this break the "clean architecture" spec
// find better way of doing it
func MapRowToColumn(rows *sql.Rows) (Column, error) {
	var t Column
	err := rows.Scan(&t.ID, &t.Name, &t.ProjectID, &t.Position)

	slog.Info("ROW TO COLUMN", err, t)

	if err != nil {
		return Column{}, fmt.Errorf("Error error scanning COLUMN row: %w", err)

	}
	return t, nil
}
