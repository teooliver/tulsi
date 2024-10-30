package seedDb

import (
	"fmt"
	"testing"

	"github.com/teooliver/kanban/internal/repository/project"
)

func TestProjectIntoCSVString(t *testing.T) {

	projects := []project.Project{
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c5",
			Name:        "some_name",
			Description: "some_description",
			IsArchived:  false,
		},
		{
			ID:          "09a39094-4b42-4c25-96be-a0c16ee9f1c6",
			Name:        "some_name_v2",
			Description: "some_description_v2",
			IsArchived:  false,
		},
	}

	got := projectsIntoCSVString(projects)

	want := []string{
		fmt.Sprint("id,name,description,is_archived"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c5,some_name,some_description,false"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c6,some_name_v2,some_description_v2,false"),
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
