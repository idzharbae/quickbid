package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type walletReader struct {
	dbConn db.Connection
}

func NewWalletReader(dbConn db.Connection) src.WalletReaderRepo {
	return &walletReader{dbConn: dbConn}
}

func (wr *walletReader) GetByUserID(ctx context.Context, userID int) (entity.Wallet, error) {
	var wallet entity.Wallet
	err := wr.dbConn.QueryRow(ctx, `SELECT id, balance, user_id FROM wallet WHERE user_id = $1`, userID).
		Scan(&wallet.ID, &wallet.Amount, &wallet.UserID)
	if err != nil {
		return wallet, stacktrace.Propagate(err, "[walletReader][GetByUserID][QueryRow]")
	}

	return wallet, nil
}

func (wr *walletReader) WithTx(tx db.Tx) src.WalletReaderRepo {
	return NewWalletReader(tx)
}
