-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN role VARCHAR(50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN role;
-- +goose StatementEnd
