package seedDb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/column"
)

func createMultipleColumns(projectId string) []column.Column {
	return []column.Column{
		{
			ID:        uuid.New().String(),
			Name:      "Backlog",
			ProjectID: projectId,
			Position:  0,
		},
		{
			ID:        uuid.New().String(),
			Name:      "Todo",
			ProjectID: projectId,
			Position:  1,
		},
		{
			ID:        uuid.New().String(),
			Name:      "In Progress",
			ProjectID: projectId,
			Position:  2,
		},
		{
			ID:        uuid.New().String(),
			Name:      "Review",
			ProjectID: projectId,
			Position:  3,
		},
		{
			ID:        uuid.New().String(),
			Name:      "Done",
			ProjectID: projectId,
			Position:  4,
		},
	}
}

func columnsIntoCSVString(column []column.Column) []string {
	s := make([]string, 0, len(column))

	columnsCSVHeader := "id,name,position,project_id"
	s = append(s, columnsCSVHeader)

	for _, p := range column {
		result := fmt.Sprintf("%s,%s,%v,%s", p.ID, p.Name, p.Position, p.ProjectID)
		s = append(s, result)
	}

	return s
}
