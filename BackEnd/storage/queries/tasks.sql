-- name: CreateTask :one
insert into tasks (
    ID, 
    ColumnID, 
    UserID, 
    Task_Number, 
    Title, 
    Task_Description, 
    Created, 
    Completed, 
    Due, 
    Created_At, 
    Updated_At, 
    Task_Priority, 
    Task_Status, 
    Estimation
)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14
)
returning *;