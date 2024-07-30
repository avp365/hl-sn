-- +goose Up
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    text text,
    date timestamp
);

-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE posts
-- +goose StatementBegin
-- +goose StatementEnd
