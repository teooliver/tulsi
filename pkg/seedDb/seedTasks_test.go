package seedDb

import (
	"fmt"
	"slices"
	"testing"

	"github.com/teooliver/kanban/internal/repository/task"
)

func TestTaskIntoCSVString(t *testing.T) {
	statusID := "09a39094-4b42-4c25-96be-a0c16ee9f1c6"
	userID := "09a39094-4b42-4c25-96be-a0c16ee9f1c7"

	tasks := []task.Task{
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title",
			Description: "some_description",
			Color:       "some_color",
			StatusID:    &statusID,
			UserID:      &userID,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title_v2",
			Description: "some_description_v2",
			Color:       "some_color_v2",
			StatusID:    &statusID,
			UserID:      &userID,
		},
	}

	got := taskIntoCSVString(tasks)

	want := []string{
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title,some_color,some_description,%v,%v", statusID, userID),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title_v2,some_color_v2,some_description_v2,%v,%v", statusID, userID),
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %q, \n wanted %q", got[0], want[0])
	}
}
