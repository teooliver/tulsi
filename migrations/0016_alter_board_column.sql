-- +goose up
ALTER TABLE IF EXISTS "board_column"
RENAME TO "project_column";

ALTER TABLE IF EXISTS "board_column"
RENAME COLUMN board_id to project_id;
