-- name: CreateBoard :one
insert into boards (id, title) 
values ($1, $2)
returning *;