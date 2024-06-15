package seedDb

import (
	"github.com/teooliver/kanban/internal/repository/status"
)

func CreateMultipleStatus() []status.StatusForCreate {

	return []status.StatusForCreate{
		{
			Name: "Backlog",
		},
		{
			Name: "Todo",
		},
		{
			Name: "In Progress",
		},
		{
			Name: "Review",
		},
		{
			Name: "Done",
		},
	}
}
