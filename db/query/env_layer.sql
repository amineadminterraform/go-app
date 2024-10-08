

-- name: CreateEnvLayer :one
INSERT INTO env_layer (
    environment_id,
    s3_path,
    process_id,
    current_request_id
) VALUES (
    $1, $2, $3, $4
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
  set s3_path = $2,
  process_id = $3,
  current_request_id = $4,
  updated_at = now()
  WHERE id = $1;

-- name: DeleteEnvLayer :exec
DELETE FROM env_layer
WHERE id = $1;