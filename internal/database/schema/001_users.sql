-- +goose Up

CREATE TABLE users
(
  id         UUID PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  name TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;