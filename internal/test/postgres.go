package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	defaultTimeout = 5000 * time.Second
)

const (
	ENV_DSN = "DSN"
	// Note: this config comes from compose.yml.
	defaultTestDsn = "host=localhost port=5432 user=db_user dbname=kanban-go password=12345 sslmode=disable"
)

func TestDB(ctx context.Context, t *testing.T) (*sql.DB, error) {
	t.Helper()
	db, err := sql.Open("pgx", defaultTestDsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to db", err)
		return db, fmt.Errorf("error connecting to db: %w", err)
	}

	log.Println("Database connection established")

	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(20)
	cleanDB(ctx, t, db)

	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	return db, nil
}

func cleanDB(ctx context.Context, t *testing.T, db *sql.DB) {
	t.Helper()

	tables, err := listAllTables(ctx, db)
	require.NoError(t, err)

	// Disable foreign key checks
	_, err = db.ExecContext(ctx, "SET session_replication_role = replica;")
	require.NoError(t, err)

	for _, table := range tables {
		_, err := db.ExecContext(ctx, "DELETE FROM "+table)
		require.NoError(t, err)
	}

	_, err = db.ExecContext(ctx, "SET session_replication_role = DEFAULT;")
	require.NoError(t, err)
}

func listAllTables(ctx context.Context, db *sql.DB) (_ []string, err error) {
	var tables []string
	rows, err := db.QueryContext(ctx, `
		SELECT tablename
		FROM pg_tables
		WHERE schemaname = 'public' AND tablename != 'goose_db_version';
	`)
	if err != nil {
		return nil, err
	}
	defer func() {
		// JOIN ERRORS
	}()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}
