// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: columns.sql

package database

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createColumn = `-- name: CreateColumn :one
insert into columns (id, boardId, title) 
values ($1, $2, $3)
returning id, boardid, title
`

type CreateColumnParams struct {
	ID      uuid.UUID
	Boardid uuid.UUID
	Title   string
}

func (q *Queries) CreateColumn(ctx context.Context, arg CreateColumnParams) (Column, error) {
	row := q.db.QueryRowContext(ctx, createColumn, arg.ID, arg.Boardid, arg.Title)
	var i Column
	err := row.Scan(&i.ID, &i.Boardid, &i.Title)
	return i, err
}

const getBoardIdColums = `-- name: GetBoardIdColums :many
select id, boardid, title from columns where boardId=$1
`

func (q *Queries) GetBoardIdColums(ctx context.Context, boardid uuid.UUID) ([]Column, error) {
	rows, err := q.db.QueryContext(ctx, getBoardIdColums, boardid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Column
	for rows.Next() {
		var i Column
		if err := rows.Scan(&i.ID, &i.Boardid, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBoardIdData = `-- name: GetBoardIdData :many
SELECT
    json_build_object(
        'column_id', c.id,
        'column_title', c.title,
        'tasks', CASE WHEN COUNT(t.id) > 0 THEN json_agg(
            json_build_object(
                'task_id', t.id,
                'task_number', t.task_number,
                'task_title', t.title,
                'task_description', t.task_description,
                'created', t.created,
                'completed', t.completed,
                'due', t.due,
                'created_at', t.created_at,
                'updated_at', t.updated_at,
                'task_priority', t.task_priority,
                'task_status', t.task_status,
                'estimation', t.estimation,
                'user_id', t.userid
            )
        ) ELSE '[]'::json END) AS column_with_tasks
FROM
    columns c
LEFT JOIN
    tasks t ON c.id = t.columnid
JOIN
    boards b ON c.boardid = b.id
WHERE
    b.id = $1
GROUP BY
    c.id, c.title
`

func (q *Queries) GetBoardIdData(ctx context.Context, id uuid.UUID) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getBoardIdData, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var column_with_tasks json.RawMessage
		if err := rows.Scan(&column_with_tasks); err != nil {
			return nil, err
		}
		items = append(items, column_with_tasks)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
