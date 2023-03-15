-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth_user 
(
    id serial NOT NULL PRIMARY KEY,
    login varchar(30) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar(128) NOT NULL,
    phone_number varchar(20) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth_user;
-- +goose StatementEnd
