package user

import (
	"context"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/pkg/testhelpers"
)

type TaskRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *PostgresRepository
	ctx         context.Context
}

func (suite *TaskRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	postgresConfig := config.PostgresConfig{
		DSN: pgContainer.ConnectionString,
	}

	db, err := testhelpers.InitPostgres(suite.ctx, &postgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer
	taskRepo := NewPostgres(db)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = taskRepo
}

func (suite *TaskRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *TaskRepoTestSuite) TestUserRepo() {
	t := suite.T()

	// Create Task 01
	user01ID, err := suite.repository.CreateUser(suite.ctx, UserForCreate{
		Email:     "some_user@email.com",
		FirstName: "first_name",
		LastName:  "last_name",
	})
	assert.NoError(t, err)
	assert.NotNil(t, user01ID)
}

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
