-- +goose up
ALTER TABLE status
DROP COLUMN sprint_id;
