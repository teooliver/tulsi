-- +goose up
ALTER TABLE task
DROP COLUMN status;
