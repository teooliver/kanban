package task

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/teooliver/kanban/internal/bootstrap"
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

	testConfig := config.Config{
		Postgres: postgresConfig,
	}

	deps, err := bootstrap.Deps(suite.ctx, &testConfig)
	if err != nil {
		log.Fatal("Error bootstraping application: %w", err)
		panic("error bootstraping application")
	}

	//Question: Why TaskRepo is *invalid type instead of *task.PostgresRepository
	var taskRepo = deps.Repos.TaskRepo

	suite.pgContainer = pgContainer
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

func (suite *TaskRepoTestSuite) TestCreateTask() {
	t := suite.T()

	id, err := suite.repository.CreateTask(suite.ctx, TaskForCreate{
		Title:       "some title",
		Description: "some description",
		Color:       "some color",
	})
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
