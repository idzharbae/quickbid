package pgx

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -destination=pgxmock/pgx_mock.go -package=pgxmock -source=pgxinterface.go

type Pgx interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}
