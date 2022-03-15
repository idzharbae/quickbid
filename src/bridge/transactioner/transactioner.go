package transactioner

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"
)

//go:generate mockgen -destination=txnermock/txn_mock.go -package=txnermock -source=transactioner.go
type Transactioner interface {
	DoWithTx(ctx context.Context, f func(context.Context, db.Tx) error) error
}
