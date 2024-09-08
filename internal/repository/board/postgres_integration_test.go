package board

import (
	"context"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/teooliver/kanban/internal/config"
	"github.com/teooliver/kanban/pkg/postgresutils"
	"github.com/teooliver/kanban/pkg/testhelpers"
)

type BoardRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *PostgresRepository
	ctx         context.Context
}

func (suite *BoardRepoTestSuite) SetupSuite() {
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
	boardRepo := NewPostgres(db)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = boardRepo
}

func (suite *BoardRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *BoardRepoTestSuite) TestBoardRepo() {
	t := suite.T()

	// Create Board 01
	board01ID, err := suite.repository.CreateBoard(suite.ctx, BoardForCreate{
		Name: "board_01",
	})
	assert.NoError(t, err)
	assert.NotNil(t, board01ID)

	// Update Board 01
	err = suite.repository.UpdateBoard(suite.ctx, board01ID, BoardForUpdate{
		Name: "updated_board_01",
	})
	assert.NoError(t, err)

	// Create Board 02
	id, err := suite.repository.CreateBoard(suite.ctx, BoardForCreate{
		Name: "board_02",
	})
	assert.NoError(t, err)
	assert.NotNil(t, id)

	// Delete Board 02
	id, err = suite.repository.DeleteBoard(suite.ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, id)

	// Check final Board Repo state:
	actual, err := suite.repository.ListAllBoards(suite.ctx, &postgresutils.PageRequest{
		Page: 0,
		Size: 0, // All items
	})

	assert.Equal(t, 1, int(actual.TotalElements))
	assert.Equal(t, 1, int(actual.TotalPages))
}

func TestBoardRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BoardRepoTestSuite))
}
