package db

import "context"

//go:generate mockgen -destination=dbmock/db_mock.go -package=dbmock -source=db.go

type Result interface {
	RowsAffected() (int64, error)
}

type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Connection interface {
	Begin(ctx context.Context) (Tx, error)

	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
	Exec(ctx context.Context, query string, args ...interface{}) (Result, error)
}

type Stmt interface {
	Connection

	Close() error
}

type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Connection
}
