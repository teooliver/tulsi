package task

import (
	"context"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/teooliver/tulsi/internal/config"
	"github.com/teooliver/tulsi/pkg/postgresutils"
	"github.com/teooliver/tulsi/pkg/testhelpers"
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

func (suite *TaskRepoTestSuite) TestTaskRepo() {
	t := suite.T()

	// Create Task 01
	task01ID, err := suite.repository.CreateTask(suite.ctx, TaskForCreate{
		Title:       "some title 01",
		Description: "some description 01",
		Color:       "some color 01",
	})
	assert.NoError(t, err)
	assert.NotNil(t, task01ID)

	// Update Task 01
	err = suite.repository.UpdateTask(suite.ctx, task01ID, TaskForUpdate{
		Title:       "updated title",
		Description: "updated description",
		Color:       "updated color",
	})
	assert.NoError(t, err)

	// Create Task 02
	id, err := suite.repository.CreateTask(suite.ctx, TaskForCreate{
		Title:       "some title 02",
		Description: "some description 02",
		Color:       "some color 02",
	})
	assert.NoError(t, err)
	assert.NotNil(t, id)

	// Delete Task 02
	id, err = suite.repository.DeleteTask(suite.ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, id)

	// Check final Task Repo state:
	actual, err := suite.repository.ListAllTasks(suite.ctx, &postgresutils.PageRequest{
		Page: 0,
		Size: 0, // All items
	})

	assert.Equal(t, 51, int(actual.TotalElements))
	assert.Equal(t, 1, int(actual.TotalPages))
}

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
