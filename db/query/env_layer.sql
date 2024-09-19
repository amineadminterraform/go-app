

-- name: CreateEnvLayer :one
INSERT INTO env_layer (
    environment_id,
    s3_path,
    template_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetEnvLayer :one
SELECT * FROM env_layer
WHERE id = $1 LIMIT 1;

-- name: ListEnvLayers :many
SELECT * FROM env_layer
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateEnvLayer :exec
UPDATE env_layer
  set environment_id = $2,
  s3_path = $3,
  template_id = $4
WHERE id = $1;

-- name: DeleteEnvLayer :exec
DELETE FROM env_layer
WHERE id = $1;