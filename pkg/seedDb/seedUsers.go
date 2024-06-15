package seedDb

import (
	"github.com/jaswdr/faker/v2"
	"github.com/teooliver/kanban/internal/repository/user"
)

func createRandomUser() user.UserForCreate {
	fake := faker.New()

	user := user.UserForCreate{
		Email:     fake.Person().Contact().Email,
		FirstName: fake.Person().FirstName(),
		LastName:  fake.Person().LastName(),
	}

	return user
}

func CreateMultipleUsers(nbUsers int) []user.UserForCreate {
	users := make([]user.UserForCreate, nbUsers)
	user := createRandomUser()

	for i := 0; i < nbUsers; i++ {
		users = append(users, user)
	}

	return users
}
