package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/palantir/stacktrace"
)

type walletWriter struct {
	dbConn db.Connection
}

func NewWalletWriter(dbConn db.Connection) src.WalletWriterRepo {
	return &walletWriter{dbConn: dbConn}
}

func (wr *walletWriter) DeductWallet(ctx context.Context, id, deduction int) error {
	_, err := wr.dbConn.Exec(ctx, `UPDATE wallet SET balance = balance - $1 WHERE id = $2`,
		deduction, id)
	if err != nil {
		return stacktrace.Propagate(err, "[walletReader][DeductWallet][QueryRow]")
	}

	return nil
}

func (wr *walletWriter) InjectWalletByUserID(ctx context.Context, userID, deduction int) error {
	_, err := wr.dbConn.Exec(ctx, `UPDATE wallet SET balance = balance + $1 WHERE user_id = $2`,
		deduction, userID)
	if err != nil {
		return stacktrace.Propagate(err, "[walletReader][InjectWalletByUserID][QueryRow]")
	}

	return nil
}

func (wr *walletWriter) WithTx(tx db.Tx) src.WalletWriterRepo {
	return NewWalletWriter(tx)
}
