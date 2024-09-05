-- +goose up
COPY app_user FROM '/users.csv' WITH (FORMAT csv);
COPY status FROM '/status.csv' WITH (FORMAT csv);
COPY task FROM '/tasks.csv' WITH (FORMAT csv, HEADER match);
