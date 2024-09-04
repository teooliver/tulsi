-- +goose up
ALTER TABLE task
ADD COLUMN sprint_id uuid REFERENCES "sprint"(id);
