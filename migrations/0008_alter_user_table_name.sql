-- +goose up
ALTER TABLE IF EXISTS "user"
RENAME TO "app_user";
