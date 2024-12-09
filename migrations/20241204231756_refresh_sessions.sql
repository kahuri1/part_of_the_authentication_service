-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS refreshSessions (
id SERIAL PRIMARY KEY,
user_uuid uuid REFERENCES users(uuid) ON DELETE CASCADE UNIQUE,
refresh_token VARCHAR(255) NOT NULL,
ip character varying(15) NOT NULL,
expires_at timestamp NOT NULL,
created_at timestamp default now() not null
);
CREATE INDEX idx_refresh_sessions_expires_at ON refreshSessions (expires_at);
CREATE INDEX idx_refresh_sessions_refresh_token ON refreshSessions (refresh_token);
CREATE INDEX idx_refresh_sessions_ip ON refreshSessions (ip);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS refreshSessions
-- +goose StatementEnd
