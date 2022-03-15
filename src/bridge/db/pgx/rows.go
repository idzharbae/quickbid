package pgx

import (
	"github.com/jackc/pgx/v4"
)

type Rows struct {
	rows pgx.Rows
}

func (r *Rows) Close() error {
	r.rows.Close()
	return nil
}
func (r *Rows) Next() bool {
	return r.rows.Next()
}
func (r *Rows) Scan(dest ...interface{}) error {
	return r.rows.Scan()
}
