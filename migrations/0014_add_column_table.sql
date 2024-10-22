-- +goose Up
CREATE TABLE IF NOT EXISTS board_column (
id uuid DEFAULT gen_random_uuid(),
name VARCHAR(200),
position VARCHAR(200),
board_id uuid REFERENCES board(id),
PRIMARY KEY(id)
);
