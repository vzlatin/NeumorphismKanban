-- name: CreateColumn :one
insert into columns (id, boardId, title) 
values ($1, $2, $3)
returning *;


-- name: GetBoardIdColums :many
select * from columns where boardId=$1;


-- name: GetBoardIdData :many
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
    c.id, c.title;

