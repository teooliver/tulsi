package deps

import (
	"context"

	"log"

	"github.com/jackc/pgx/v5"
	"github.com/teooliver/kanban/internal/config"
)

type Infra struct {
	Postgres *pgx.Conn
}

func InitInfra(ctx context.Context, cfg *config.Config) (*Infra, error) {
	// TODO: handle err
	pgClient, _ := initPostgres(context.TODO(), &cfg.Postgres)

	return &Infra{
		Postgres: pgClient,
	}, nil
}

func initPostgres(ctx context.Context, cfg *config.PostgresConfig) (*pgx.Conn, error) {
	// db, err := sql.Open("postgres", cfg.DSN)

	conn, err := pgx.Connect(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		// TODO: Better error handling
		log.Fatal("Error connecting to db")
		return conn, err
	}
	log.Println("Database connection established")

	return conn, nil
}
