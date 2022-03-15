package pgx

import "github.com/jackc/pgconn"

type Result struct {
	cmd pgconn.CommandTag
}

func (r *Result) RowsAffected() (int64, error) {
	return r.cmd.RowsAffected(), nil
}
