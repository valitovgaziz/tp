-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE INDEX,
    password VARCHAR(500) NOT NULL,
    phone VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
