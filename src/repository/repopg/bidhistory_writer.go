package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type bidHistoryWriter struct {
	dbConn db.Connection
}

func NewBidHistoryWriter(dbConn db.Connection) src.BidHistoryWriterRepo {
	return &bidHistoryWriter{dbConn: dbConn}
}

func (bw *bidHistoryWriter) Insert(ctx context.Context, bid entity.Bid) (entity.Bid, error) {
	bidHistory := bid
	err := bw.dbConn.QueryRow(ctx,
		`INSERT INTO bid_history(product_id, user_id, amount, status, bid_time) VALUES($1, $2, $3, $4, $5)
		RETURNING id`,
		bid.ProductID, bid.UserID, bid.Amount, bid.Status, bid.BidTime).Scan(&bidHistory.ID)
	if err != nil {
		return bidHistory, stacktrace.Propagate(err, "[bidWriter][Insert][QueryRow]")
	}

	return bidHistory, nil
}

func (bw *bidHistoryWriter) WithTx(tx db.Tx) src.BidHistoryWriterRepo {
	return NewBidHistoryWriter(tx)
}
