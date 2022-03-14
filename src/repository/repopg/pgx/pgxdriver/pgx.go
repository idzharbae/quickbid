package pgxdriver

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type pgxDriver struct {
	connectionString string
}

func NewPgxDriver(connectionString string) *pgxDriver {
	return &pgxDriver{
		connectionString: connectionString,
	}
}

func (p *pgxDriver) QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, error) {
	conn, err := pgx.Connect(ctx, p.connectionString)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	return conn.QueryRow(ctx, query, args...), nil
}

func (p *pgxDriver) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	conn, err := pgx.Connect(context.Background(), p.connectionString)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	ct, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return ct, nil
}
