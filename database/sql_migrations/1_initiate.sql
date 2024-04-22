-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE books (
    id BIGINT NOT NULL,
    title VARCHAR(256),
    description VARCHAR(256),
    release_year INT,
    price VARCHAR(10)
    total_page INT,
    thickness VARCHAR(6),
)

-- +migrate StatementEnd


