-- +goose Up

CREATE TABLE columns (
	id UUID NOT NULL PRIMARY KEY,
    boardId UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
	title VARCHAR(50) NOT NULL
);

-- +goose Down

DROP TABLE columns;