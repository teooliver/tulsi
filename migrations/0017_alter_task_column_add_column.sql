-- +goose up
ALTER TABLE task
ADD COLUMN column_id uuid REFERENCES "project_column"(id);
