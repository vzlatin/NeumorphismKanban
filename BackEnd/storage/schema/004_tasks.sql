-- +goose Up

CREATE TABLE tasks (
    ID UUID NOT NULL PRIMARY KEY,
    ColumnID UUID NOT NULL REFERENCES columns(id) ON DELETE CASCADE,
    UserID UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    Task_Number VARCHAR(10) NOT NULL,
    Title VARCHAR(200) NOT NULL,
    Task_Description TEXT NOT NULL,
    Created DATE NOT NULL,
    Completed DATE,
    Due DATE NOT NULL,
    Created_At DATE NOT NULL,
    Updated_At Date NOT NULL,
    Task_Priority INTEGER NOT NULL CHECK(Task_Priority >= 1 AND Task_Priority <= 4),
    Task_Status INTEGER NOT NULL CHECK(Task_Status >= 1 AND Task_Status <= 4),
    Estimation VARCHAR(10) NOT NULL
);

-- +goose Down

DROP DATABASE tasks;