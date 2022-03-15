package pgx

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

//go:generate mockgen -destination=pgxmock/pgx_mock.go -package=pgxmock -source=pgxinterface.go

type PgxDriver struct {
	pool *pgxpool.Pool
}

func NewPgxDriver(pool *pgxpool.Pool) db.Connection {
	return &PgxDriver{pool: pool}
}

func (pd *PgxDriver) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	rows, err := pd.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &Rows{rows: rows}, nil
}

func (pd *PgxDriver) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return pd.pool.QueryRow(ctx, query, args...)
}

func (pd *PgxDriver) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	cmd, err := pd.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &Result{cmd: cmd}, nil
}

func (pd *PgxDriver) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := pd.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx}, nil
}
