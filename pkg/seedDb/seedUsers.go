package seedDb

import (
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"github.com/teooliver/kanban/internal/repository/user"
)

func createFakeUser() user.User {
	fake := faker.New()

	user := user.User{
		ID:        uuid.New().String(),
		Email:     fake.Person().Contact().Email,
		FirstName: fake.Person().FirstName(),
		LastName:  fake.Person().LastName(),
	}

	return user
}

func createMultipleFakeUsers(nbUsers int) []user.User {
	users := make([]user.User, nbUsers)
	user := createFakeUser()

	for i := 0; i < nbUsers; i++ {
		users = append(users, user)
	}

	return users
}
