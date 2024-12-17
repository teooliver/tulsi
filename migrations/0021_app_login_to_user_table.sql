-- +goose up
ALTER TABLE "app_user"
ADD COLUMN "hashed_password" VARCHAR(255),
ADD COLUMN "session_token" VARCHAR(255),
ADD COLUMN "csrf_token" VARCHAR(255);
