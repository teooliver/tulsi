package project

import (
	"database/sql"
	"fmt"
)

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectToCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectToUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var allColumns = []any{
	"id",
	"name",
	"description",
}

func mapRowToProject(rows *sql.Rows) (Project, error) {
	var t Project
	err := rows.Scan(&t.ID, &t.Name, &t.Description)

	if err != nil {
		return Project{}, fmt.Errorf("Error error scanning Project row: %w", err)
	}

	return t, nil
}
