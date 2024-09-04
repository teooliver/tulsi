package seedDb

import (
	"fmt"
	"testing"

	"github.com/teooliver/kanban/internal/repository/task"
)

func TestTaskIntoCSVString(t *testing.T) {
	statusID := "09a39094-4b42-4c25-96be-a0c16ee9f1c6"
	userID := "09a39094-4b42-4c25-96be-a0c16ee9f1c7"
	// sprintID := "09a39094-4b42-4c25-96be-a0c16ee0000"

	tasks := []task.Task{
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title",
			Description: "some_description",
			Color:       "some_color",
			StatusID:    &statusID,
			UserID:      &userID,
			// SprintID:    &sprintID,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title_v2",
			Description: "some_description_v2",
			Color:       "some_color_v2",
			StatusID:    &statusID,
			UserID:      &userID,
			// SprintID:    &sprintIDINFO UPDATED ID !BADKEY
			//
			// ,
		},
	}

	got := taskIntoCSVString(tasks)

	want := []string{
		fmt.Sprint("id,title,description,color,status_id,user_id"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title,some_description,some_color,%v,%v", statusID, userID),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title_v2,some_description_v2,some_color_v2,%v,%v", statusID, userID),
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
