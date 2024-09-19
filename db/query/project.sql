-- name: CreateProject :one
INSERT INTO project (
    name,
    git_path,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProject :one
SELECT * FROM project
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateProject :exec
UPDATE project
  set name = $2,
  git_path = $3,
  description = $4
WHERE id = $1;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1;