package seedDb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/status"
)

func createFakeStatusList() []status.Status {
	return []status.Status{
		{
			ID:   uuid.New().String(),
			Name: "Backlog",
		},
		{
			ID:   uuid.New().String(),
			Name: "Todo",
		},
		{
			ID:   uuid.New().String(),
			Name: "In Progress",
		},
		{
			ID:   uuid.New().String(),
			Name: "Review",
		},
		{
			ID:   uuid.New().String(),
			Name: "Done",
		},
	}
}

func statusIntoCSVString(status []status.Status) []string {
	s := make([]string, 0, len(status))

	for _, t := range status {
		result := fmt.Sprintf("%s, %s", t.ID, t.Name)
		s = append(s, result)
	}

	return s
}
