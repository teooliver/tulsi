package seedDb

import (
	"fmt"
	"testing"

	"github.com/teooliver/kanban/internal/repository/column"
)

func TestColumnIntoCSVString(t *testing.T) {
	projectId := "09a39094-4b42-4c25-96be-a0c16ee9f1c6"

	columns := []column.Column{
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Name:        "Backlog",
			ProjectID:   projectId,
			Position:    0,
			PositionInt: 0,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c6",
			Name:        "Todo",
			ProjectID:   projectId,
			Position:    1,
			PositionInt: 1,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c7",
			Name:        "In Progress",
			ProjectID:   projectId,
			Position:    2,
			PositionInt: 2,
		},
	}

	got := columnsIntoCSVString(columns)

	want := []string{
		fmt.Sprint("id,name,position,project_id,position_int"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,Backlog,0,%s,0", projectId),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c6,Todo,1,%s,1", projectId),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c7,In Progress,2,%s,2", projectId),
	}

	if got[0] != want[0] {
		t.Errorf("got %q, \n wanted %q", got[0], want[0])
	}
	if got[1] != want[1] {
		t.Errorf("got %q, \n wanted %q", got[1], want[1])
	}
	if got[2] != want[2] {
		t.Errorf("got %q, \n wanted %q", got[2], want[2])
	}
}
