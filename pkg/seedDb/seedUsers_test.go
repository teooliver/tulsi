package seedDb

import (
	"fmt"
	"slices"
	"testing"

	"github.com/teooliver/kanban/internal/repository/user"
)

func TestUserIntoCSVString(t *testing.T) {

	users := []user.User{
		{
			ID:        "09a39094-4b42-4c25-96be-a0c16ee9f1c6",
			Email:     "email_01@email.com",
			FirstName: "first_name_01",
			LastName:  "last_name_01",
		},
		{
			ID:        "09a39094-4b42-4c25-96be-a0c16ee9f1c7",
			Email:     "email_02@email.com",
			FirstName: "first_name_02",
			LastName:  "last_name_02",
		},
	}

	got := userIntoCSVString(users)

	want := []string{
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c6, email_01@email.com, first_name_01, last_name_01"),
		fmt.Sprintf("09a39094-4b42-4c25-96be-a0c16ee9f1c7, email_02@email.com, first_name_02, last_name_02"),
	}

	if !slices.Equal(got, want) {
		t.Errorf("\n got    %q, \n wanted %q", got[0], want[0])
	}
}
