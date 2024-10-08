// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: request.sql

package db

import (
	"context"
)

const createRequest = `-- name: CreateRequest :one
INSERT INTO request (
    layer_id,
    source,
    payload,
    status
) VALUES (
    $1, $2, $3, $4
) RETURNING id, layer_id, source, payload, status, updated_at, created_at
`

type CreateRequestParams struct {
	LayerID int64  `json:"layer_id"`
	Source  string `json:"source"`
	Payload string `json:"payload"`
	Status  string `json:"status"`
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error) {
	row := q.db.QueryRow(ctx, createRequest,
		arg.LayerID,
		arg.Source,
		arg.Payload,
		arg.Status,
	)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.LayerID,
		&i.Source,
		&i.Payload,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM request
WHERE id = $1
`

func (q *Queries) DeleteRequest(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRequest, id)
	return err
}

const getRequest = `-- name: GetRequest :one
SELECT id, layer_id, source, payload, status, updated_at, created_at FROM request
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.db.QueryRow(ctx, getRequest, id)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.LayerID,
		&i.Source,
		&i.Payload,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listRequests = `-- name: ListRequests :many
SELECT id, layer_id, source, payload, status, updated_at, created_at FROM request
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRequestsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRequests(ctx context.Context, arg ListRequestsParams) ([]Request, error) {
	rows, err := q.db.Query(ctx, listRequests, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.LayerID,
			&i.Source,
			&i.Payload,
			&i.Status,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRequest = `-- name: UpdateRequest :exec
UPDATE request
  set layer_id = $2,
    source = $3,
    payload = $4,
    status= $5,
    updated_at = now()
WHERE id = $1
`

type UpdateRequestParams struct {
	ID      int64  `json:"id"`
	LayerID int64  `json:"layer_id"`
	Source  string `json:"source"`
	Payload string `json:"payload"`
	Status  string `json:"status"`
}

func (q *Queries) UpdateRequest(ctx context.Context, arg UpdateRequestParams) error {
	_, err := q.db.Exec(ctx, updateRequest,
		arg.ID,
		arg.LayerID,
		arg.Source,
		arg.Payload,
		arg.Status,
	)
	return err
}
