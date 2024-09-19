-- name: CreateProject :one
INSERT INTO project (
    name,
    git_path,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;
