package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserForCreate struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserForUpdate struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var allColumns = []any{
	"id",
	"email",
	"first_name",
	"last_name",
}

func mapRowToStatus(rows *sql.Rows) (User, error) {
	var u User
	err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName)

	if err != nil {
		return User{}, fmt.Errorf("Error error scanning User row: %w", err)
	}
	return u, nil
}
