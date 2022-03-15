package pgx

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"

	"github.com/jackc/pgx/v4"
)

type Tx struct {
	tx pgx.Tx
}

func NewTransaction(tx pgx.Tx) db.Tx {
	return &Tx{tx: tx}
}

func (tx *Tx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return tx.tx.Rollback(ctx)
}

func (tx *Tx) Begin(ctx context.Context) (db.Tx, error) {
	newTx, err := tx.tx.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{tx: newTx}, nil
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	rows, err := tx.tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &Rows{rows: rows}, nil
}

func (tx *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return tx.tx.QueryRow(ctx, query, args...)
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	res, err := tx.tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &Result{cmd: res}, nil
}
