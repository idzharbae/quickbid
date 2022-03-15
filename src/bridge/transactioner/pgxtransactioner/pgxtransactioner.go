package pgxtransactioner

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/bridge/db/pgx"
	"github.com/idzharbae/quickbid/src/bridge/transactioner"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgxTransactioner struct {
	pool *pgxpool.Pool
}

func NewPgxTransactioner(pool *pgxpool.Pool) transactioner.Transactioner {
	return &PgxTransactioner{pool: pool}
}

func (txn *PgxTransactioner) DoWithTx(ctx context.Context, f func(ctx context.Context, tx db.Tx) error) error {
	tx, err := txn.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	err = f(ctx, pgx.NewTransaction(tx))
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
