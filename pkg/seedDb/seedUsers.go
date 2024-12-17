package seedDb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/user"
)

func createFakeUser() user.User {
	login := user.Login{HashedPassword: "", SessionToken: "", CSRFToken: ""}

	user := user.User{
		ID:        uuid.New().String(),
		Email:     fake.Person().Contact().Email,
		FirstName: fake.Person().FirstName(),
		LastName:  fake.Person().LastName(),
		Login:     login,
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

	usersCSVHeader := "id,email,first_name,last_name,hashed_password,session_token,csrf_token"
	s = append(s, usersCSVHeader)

	for _, t := range users {
		result := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
			t.ID, t.Email, t.FirstName, t.LastName, t.Login.HashedPassword, t.Login.SessionToken, t.Login.CSRFToken)
		s = append(s, result)
	}

	return s
}
