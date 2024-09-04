-- +goose up
ALTER TABLE task
DROP COLUMN sprint_id;
