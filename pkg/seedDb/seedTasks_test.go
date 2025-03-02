package seedDb

import (
	"fmt"
	"testing"

	"github.com/teooliver/tulsi/internal/repository/task"
)

func TestTaskIntoCSVString(t *testing.T) {
	statusID := "09a39094-4b42-4c25-96be-a0c16ee9f1c6"
	userID := "09a39094-4b42-4c25-96be-a0c16ee9f1c7"
	columnID := "09a39094-4b42-4c25-96be-a0c16ee9f1c8"

	tasks := []task.Task{
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title",
			Description: "some_description",
			Color:       "some_color",
			StatusID:    &statusID,
			UserID:      &userID,
			ColumnID:    columnID,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Title:       "some_title_v2",
			Description: "some_description_v2",
			Color:       "some_color_v2",
			StatusID:    &statusID,
			UserID:      &userID,
			ColumnID:    columnID,
		},
	}

	got := taskIntoCSVString(tasks)

	want := []string{
		fmt.Sprint("id,title,description,color,user_id,status_id,column_id"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title,some_description,some_color,%v,%v,%v", userID, statusID, columnID),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_title_v2,some_description_v2,some_color_v2,%v,%v,%v", userID, statusID, columnID),
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
