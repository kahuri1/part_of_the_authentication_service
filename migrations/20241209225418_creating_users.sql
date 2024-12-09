-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO users (uuid, name, password_hash, email, created_at, updated_at) VALUES
('dff723ba-4da7-4c55-9f07-27121ec53386', 'Ivan',
 '123123123123',
 'Kahuri1@github.com', NOW(), NOW()),

('dff723ba-4da7-4c55-9f07-27121ec53387', 'Oleg',
 '123123123123123123',
 'Kahuri11@github.com', NOW(), NOW());
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
