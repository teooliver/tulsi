-- +goose up
ALTER TABLE project
ADD COLUMN description VARCHAR(200);
