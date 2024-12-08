-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS refreshSessions (
id SERIAL PRIMARY KEY,
user_uuid uuid REFERENCES users(uuid) ON DELETE CASCADE,
refresh_token uuid NOT NULL,
ip character varying(15) NOT NULL,
expires_at timestamp NOT NULL,
created_at timestamp default now() not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS refreshSessions
-- +goose StatementEnd
