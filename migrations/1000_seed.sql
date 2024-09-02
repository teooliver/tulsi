COPY app_user FROM '/users.csv' WITH (FORMAT csv);
COPY status FROM '/status.csv' WITH (FORMAT csv);
COPY tasks FROM '/tasks.csv' WITH (FORMAT csv);
