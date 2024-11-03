-- +goose up
ALTER TABLE project_column
ALTER COLUMN position TYPE INT;
