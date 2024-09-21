-- name: CreateRequest :one
INSERT INTO request (
    layer_id,
    environment_id,
    payload,
    status
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetRequest :one
SELECT * FROM request
WHERE id = $1 LIMIT 1;

-- name: ListRequests :many
SELECT * FROM request
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateRequest :exec
UPDATE request
  set layer_id = $2,
    environment_id = $3,
    payload = $4,
    status= $5,
  updated_at = now()
WHERE id = $1;

-- name: DeleteRequest :exec
DELETE FROM request
WHERE id = $1;