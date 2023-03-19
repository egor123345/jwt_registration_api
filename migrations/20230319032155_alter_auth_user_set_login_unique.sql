-- +goose Up
-- +goose StatementBegin
ALTER TABLE auth_user
ADD CONSTRAINT auth_user_unique_login UNIQUE (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE auth_user
DROP CONSTRAINT auth_user_unique_login;
-- +goose StatementEnd
