-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    Id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50),
    password VARCHAR(1000) NOT NULL,
    phone VARCHAR(30) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
