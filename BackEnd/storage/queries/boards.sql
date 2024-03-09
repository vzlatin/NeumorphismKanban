-- name: CreateBoard :one
insert into boards (id, title) 
values ($1, $2)
returning *;

-- name: GetAllBoards :many
select * from boards;