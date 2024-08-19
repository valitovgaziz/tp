-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(500) NOT NULL,
    phone VARCHAR(50) NOT NULL UNIQUE
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE INDEX email_index ON users (email);
CREATE INDEX phone_index ON users (phone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
