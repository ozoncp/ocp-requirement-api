-- +goose Up
-- +goose StatementBegin
CREATE TABLE requirement
(
    id          SERIAL     PRIMARY KEY,
    user_id     INTEGER     NOT NULL,
    text        TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE requirement;
-- +goose StatementEnd
