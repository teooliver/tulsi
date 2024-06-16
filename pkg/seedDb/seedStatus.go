package seedDb

import (
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
