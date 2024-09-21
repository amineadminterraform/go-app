-- name: CreateArgoWorkflow :one
INSERT INTO argo_workflow (
    name,
    path,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetArgoWorkflow :one
SELECT * FROM argo_workflow
WHERE id = $1 LIMIT 1;

-- name: ListArgoWorkflows :many
SELECT * FROM argo_workflow
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateArgoWorkflow :exec
UPDATE argo_workflow
  set name = $2,
    path = $3,
    description =$4,
  updated_at = now()
WHERE id = $1;

-- name: DeleteArgoWorkflow :exec
DELETE FROM argo_workflow
WHERE id = $1;