package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/c2fo/testify/assert"
	"github.com/c2fo/testify/require"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/teooliver/kanban/internal/repository/user"
	"github.com/teooliver/kanban/internal/test"
)

func TestIntegration_PostgresRepository(t *testing.T) {
	createUser := func(
		ctx context.Context,
		t *testing.T,
		repo *user.PostgresRepository,
		u user.UserForCreate,
	) user.User {
		t.Helper()
		id, err := repo.CreateUser(ctx, u)
		if err != nil {
			require.NoError(t, err)
		}

		return user.User{
			ID:        id,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	}

	t.Run("Create User", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
		t.Cleanup(cancel)
		db, _ := test.DB(ctx, t)
		repo := user.NewPostgres(db)

		userToCreate := user.UserForCreate{
			Email:     "name@email.com",
			FirstName: "first_name",
			LastName:  "last_name",
		}

		userFromRepo := createUser(
			ctx,
			t,
			repo,
			userToCreate,
		)

		assert.Equal(t, userToCreate, userFromRepo)
	})

}
