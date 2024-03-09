-- +goose Up

CREATE TABLE users (
    ID UUID NOT NULL PRIMARY KEY,
    First_Name VARCHAR(30) NOT NULL,
    Last_Name VARCHAR(30) NOT NULL,
    Email VARCHAR(30) NOT NULL,
    Password TEXT NOT NULL,
    Profile_Image BYTEA
);

-- +goose Down

DROP DATABASE users;