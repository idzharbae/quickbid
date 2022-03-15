package app

import (
	"github.com/idzharbae/quickbid/src/bridge/transactioner"
	"github.com/idzharbae/quickbid/src/bridge/transactioner/pgxtransactioner"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Bridges struct {
	transactioner transactioner.Transactioner
}

func newBridges(pool *pgxpool.Pool) *Bridges {
	txner := pgxtransactioner.NewPgxTransactioner(pool)

	return &Bridges{
		transactioner: txner,
	}
}
