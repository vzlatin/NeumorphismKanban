-- +goose Up

CREATE TABLE boards (
	id UUID NOT NULL PRIMARY KEY,
	title VARCHAR(50) NOT NULL
);

-- +goose Down

DROP TABLE boards;