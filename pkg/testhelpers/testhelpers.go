package testhelpers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/teooliver/kanban/internal/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get path")
	}

	migrationFilesPath, err := filepath.Glob(filepath.Join(filepath.Dir(path), "..", "..", "migrations", "*.sql"))

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		// migrations
		postgres.WithInitScripts(migrationFilesPath...),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

// Todo: This is being duplicated from the deps module, but had to be done to avoid compiler error (cycles and using internal packages)
func InitPostgres(ctx context.Context, cfg *config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		// TODO?
		// db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to db")
		return db, fmt.Errorf("error connecting to db: %w", err)
	}

	log.Println("Database connection established")

	return db, nil
}
