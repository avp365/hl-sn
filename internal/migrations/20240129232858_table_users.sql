-- +goose Up
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    first_name text,
    second_name text,
    birthdate date,
    biography text,
    city text,
    password text
);

-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE post;
-- +goose StatementBegin
-- +goose StatementEnd
