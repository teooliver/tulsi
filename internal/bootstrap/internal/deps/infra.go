package deps

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/teooliver/kanban/internal/config"
)

type Infra struct {
	Postgres *sql.DB
}

func InitInfra(ctx context.Context, cfg *config.Config) (*Infra, error) {
	pgClient, err := initPostgres(ctx, &cfg.Postgres)
	if err != nil {
		panic(err)
	}

	return &Infra{
		Postgres: pgClient,
	}, nil
}

func initPostgres(ctx context.Context, cfg *config.PostgresConfig) (*sql.DB, error) {
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

// func initPostgres(ctx context.Context, cfg *config.PostgresConfig) (*sql.DB, error) {
// 	db, err := sql.Open("pgx", cfg.DSN)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error connecting to db")
// 		return db, fmt.Errorf("error connecting to db: %w", err)
// 	}

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			db.Close()
// 			fmt.Println("Connection Closed")
// 			return db, nil
// 		default:
// 			log.Println("Database connection established")
// 			return db, nil
// 		}
// 	}
// }
