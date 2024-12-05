-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS refreshSessions (
id SERIAL PRIMARY KEY,
userId uuid REFERENCES users(id) ON DELETE CASCADE,
refreshToken uuid NOT NULL,
ip character varying(15) NOT NULL,
expiresIn bigint NOT NULL,
createdAt timestamp default now() not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS refreshSessions
-- +goose StatementEnd
