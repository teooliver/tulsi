package seedDb

import (
	"fmt"
	"slices"
	"testing"

	"github.com/teooliver/kanban/internal/repository/status"
)

func TestStatusIntoCSVString(t *testing.T) {
	status := []status.Status{
		{
			ID:   "09a39094-4b42-4c25-96be-a0c16ee9f1c6",
			Name: "Backlog",
		},
		{
			ID:   "09a39094-4b42-4c25-96be-a0c16ee9f1c7",
			Name: "Todo",
		},
	}

	got := statusIntoCSVString(status)

	want := []string{
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c6, Backlog"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c7, Todo"),
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %q, \n wanted %q", got[0], want[0])
	}
}
