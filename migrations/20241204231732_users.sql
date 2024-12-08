-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY,
uuid UUID NOT NULL UNIQUE,
name VARCHAR(64) NOT NULL UNIQUE,
password_hash VARCHAR(64) NOT NULL,
email VARCHAR(64) NOT NULL UNIQUE,
created_at timestamp default now() not null,
updated_at timestamp default now() not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users
-- +goose StatementEnd
