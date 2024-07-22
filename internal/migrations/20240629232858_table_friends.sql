-- +goose Up
CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    id_user_1 integer NOT NULL,
    id_user_2 integer NOT NULL,
    date_add timestamp
);

-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE friends;
-- +goose StatementBegin
-- +goose StatementEnd
