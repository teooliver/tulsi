package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/teooliver/kanban/internal/config"

	"github.com/joho/godotenv"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "db_user"
	dbname = "kanban-go"
)

func Config(filename string) (*config.Config, error) {
	err := godotenv.Load(filename)
	if err != nil {
		// should ignore error since vars can be set in k8s
		log.Fatal("Error loading .env file")
		// log.Error(ctx, "failed to load env files", err)
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, postgresPassword, dbname)
	postgresConfig := config.PostgresConfig{
		DSN: psqlconn,
	}

	var cfg config.Config = config.Config{
		Postgres: postgresConfig,
	}

	return &cfg, nil
}
