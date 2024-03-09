// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tasks.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
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
returning id, columnid, userid, task_number, title, task_description, created, completed, due, created_at, updated_at, task_priority, task_status, estimation
`

type CreateTaskParams struct {
	ID              uuid.UUID
	Columnid        uuid.UUID
	Userid          uuid.UUID
	TaskNumber      string
	Title           string
	TaskDescription string
	Created         time.Time
	Completed       sql.NullTime
	Due             time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TaskPriority    int32
	TaskStatus      int32
	Estimation      string
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.Columnid,
		arg.Userid,
		arg.TaskNumber,
		arg.Title,
		arg.TaskDescription,
		arg.Created,
		arg.Completed,
		arg.Due,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.TaskPriority,
		arg.TaskStatus,
		arg.Estimation,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Columnid,
		&i.Userid,
		&i.TaskNumber,
		&i.Title,
		&i.TaskDescription,
		&i.Created,
		&i.Completed,
		&i.Due,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TaskPriority,
		&i.TaskStatus,
		&i.Estimation,
	)
	return i, err
}
