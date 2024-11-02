-- +goose up
ALTER TABLE project
ADD COLUMN is_archived BOOLEAN;
