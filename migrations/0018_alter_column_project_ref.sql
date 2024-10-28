-- +goose up
ALTER TABLE project_column
DROP COLUMN project_id,
ADD COLUMN project_id uuid REFERENCES "project"(id);
