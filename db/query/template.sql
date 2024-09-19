
CREATE TABLE "template" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "path" VARCHAR NOT NULL
);



-- name: CreateTemplate :one
INSERT INTO template (
    name,
    path
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetTemplate :one
SELECT * FROM template
WHERE id = $1 LIMIT 1;

-- name: ListTemplates :many
SELECT * FROM template
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateTemplate :exec
UPDATE template
  set name = $2,
  path = $3
WHERE id = $1;

-- name: DeleteTemplate :exec
DELETE FROM project_environment
WHERE id = $1;