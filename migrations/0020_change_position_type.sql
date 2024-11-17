-- +goose up
ALTER TABLE project_column
ADD COLUMN position_int SMALLINT;
