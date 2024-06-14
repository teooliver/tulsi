package status

import (
	"database/sql"
)

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StatusForCreate struct {
	Name string `json:"name"`
}

type StatusForUpdate struct {
	Name string `json:"name"`
}

var allColumns = []any{
	"id",
	"name",
}

func mapRowToTask(rows *sql.Rows) (Status, error) {
	var s Status
	err := rows.Scan(&s.ID, &s.Name)

	if err != nil {
		return Status{}, err
	}
	return s, nil
}
