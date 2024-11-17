package seedDb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/column"
)

func createMultipleColumns(projectId string) []column.Column {
	return []column.Column{
		{
			ID:          uuid.New().String(),
			Name:        "Backlog",
			ProjectID:   projectId,
			Position:    0,
			PositionInt: 0,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Todo",
			ProjectID:   projectId,
			Position:    1,
			PositionInt: 1,
		},
		{
			ID:          uuid.New().String(),
			Name:        "In Progress",
			ProjectID:   projectId,
			Position:    2,
			PositionInt: 2,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Review",
			ProjectID:   projectId,
			Position:    3,
			PositionInt: 3,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Done",
			ProjectID:   projectId,
			Position:    4,
			PositionInt: 4,
		},
	}
}

func columnsIntoCSVString(column []column.Column) []string {
	s := make([]string, 0, len(column))

	columnsCSVHeader := "id,name,position,project_id,position_int"
	s = append(s, columnsCSVHeader)

	for _, p := range column {
		result := fmt.Sprintf("%s,%s,%v,%s,%v", p.ID, p.Name, p.Position, p.ProjectID, p.PositionInt)
		s = append(s, result)
	}

	return s
}
