package seedDb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/user"
)

func createFakeUser() user.User {
	user := user.User{
		ID:        uuid.New().String(),
		Email:     fake.Person().Contact().Email,
		FirstName: fake.Person().FirstName(),
		LastName:  fake.Person().LastName(),
	}

	return user
}

func createMultipleFakeUsers(nbUsers int) []user.User {
	users := make([]user.User, 0, nbUsers)

	for i := 0; i < nbUsers; i++ {
		user := createFakeUser()
		users = append(users, user)
	}

	return users
}

func userIntoCSVString(users []user.User) []string {
	s := make([]string, 0, len(users))

	for _, t := range users {
		result := fmt.Sprintf("%s, %s, %s, %s", t.ID, t.Email, t.FirstName, t.LastName)
		s = append(s, result)
	}

	return s
}
