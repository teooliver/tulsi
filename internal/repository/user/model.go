package user

import (
	"database/sql"
	"fmt"
)

type Login struct {
	HashedPassword string `json:"hashed_password"`
	SessionToken   string `json:"session_token"`
	CSRFToken      string `json:"csrf_token"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     Login  `json:"login"`
}

type UserForCreate struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UserForUpdate struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     Login  `json:"login"`
}

var allColumns = []any{
	"id",
	"email",
	"first_name",
	"last_name",
	"hashed_password",
	"session_token",
	"csrf_token",
}

func mapRowToUser(rows *sql.Rows) (User, error) {
	var u User
	err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Login.HashedPassword, &u.Login.SessionToken, &u.Login.CSRFToken)

	if err != nil {
		return User{}, fmt.Errorf("Error error scanning User row: %w", err)
	}
	return u, nil
}
