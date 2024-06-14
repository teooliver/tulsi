package deps

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/teooliver/kanban/internal/config"
)

type Infra struct {
	Postgres *sql.DB
}

func InitInfra(ctx context.Context, cfg *config.Config) (*Infra, error) {
	// TODO: handle err
	pgClient, _ := initPostgres(context.TODO(), &cfg.Postgres)

	return &Infra{
		Postgres: pgClient,
	}, nil
}

func initPostgres(ctx context.Context, cfg *config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		// TODO: Better error handling
		log.Fatal("Error connecting to db")
		return db, err
	}
	log.Println("Database connection established")

	return db, nil
}
