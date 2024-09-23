package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// Store provides all functions to executes queries and transactions

type Store interface {
	Querier
}

// Store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *pgx.Conn
}

// New Store creates a new Begin(ctx context.Context) (pgx.Tx, error)store
func NewStore(db *pgx.Conn) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}

}

// execTX executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)

}

// StatusTX contains the input parameters of the status changement
type StatusTxParams struct {
	RequestID int64  `json:"envlayer_id"`
	Status    string `json:"status"`
}
type StatusTXResult struct {
	RequestID        Request  `json:"envlayer_id"`
	Status           Request  `json:"status"`
	CurrentRequestID EnvLayer `json:"current_request_id"`
}

// StatusTx updates the status of a request
// it will look for the full entries for the layer and the request and it will update the current request in the layer table according the status values
func (store *SQLStore) StatusTx(ctx context.Context, arg StatusTxParams) (StatusTXResult, error) {
	var result StatusTXResult
	err := store.execTx(ctx, func(q *Queries) error {
		full_request, err1 := q.GetRequest(ctx, arg.RequestID)
		if err1 != nil {
			return err1
		}
		full_layer, err1 := q.GetEnvLayer(ctx, full_request.EnvironmentID)
		if err1 != nil {
			return err1
		}
		if arg.Status == "Error" || arg.Status == "Finished" {
			err := q.UpdateEnvLayer(ctx, UpdateEnvLayerParams{
				ID:               full_request.EnvironmentID,
				S3Path:           full_layer.S3Path,
				ProcessID:        full_layer.ProcessID,
				CurrentRequestID: pgtype.Int8{},
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	return result, err

}
