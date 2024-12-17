-- +goose up
COPY app_user
FROM
    '/users.csv'
WITH
    (FORMAT csv, HEADER match);

COPY status
FROM
    '/status.csv'
WITH
    (FORMAT csv);

COPY project
FROM
    '/projects.csv'
WITH
    (FORMAT csv, HEADER match);

COPY project_column
FROM
    '/columns.csv'
WITH
    (FORMAT csv, HEADER match);

COPY task
FROM
    '/tasks.csv'
WITH
    (FORMAT csv, HEADER match);
