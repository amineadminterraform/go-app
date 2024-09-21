-- name: CreateProcess :one
INSERT INTO process (
    argo_id,
    name,
    template_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProcess :one
SELECT * FROM process
WHERE id = $1 LIMIT 1;

-- name: ListProcesss :many
SELECT * FROM process
ORDER BY id
LIMIT $1
OFFSET $2
;
-- name: UpdateProcess :exec
UPDATE process
  set argo_id = $2,
  name= $3,
  template_id = $4,
  updated_at = now()
WHERE id = $1;

-- name: DeleteProcess :exec
DELETE FROM process
WHERE id = $1;