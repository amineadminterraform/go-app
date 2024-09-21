

-- name: CreateProjectEnvironment :one
INSERT INTO project_environment (
    git_branch,
    project_id,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProjectEnvironment :one
SELECT * FROM project_environment
WHERE id = $1 LIMIT 1;

-- name: ListProjectEnvironments :many
SELECT * FROM project_environment
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateProjectEnvironment :exec
UPDATE project_environment
  set git_branch = $2,
  project_id = $3,
  description = $4,
  updated_at = now()
WHERE id = $1;

-- name: DeleteProjectEnvironment :exec
DELETE FROM project_environment
WHERE id = $1;