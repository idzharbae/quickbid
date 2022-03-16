package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type bidWriter struct {
	dbConn db.Connection
}

func NewBidWriter(dbConn db.Connection) src.BidWriterRepo {
	return &bidWriter{dbConn: dbConn}
}

func (bw *bidWriter) Insert(ctx context.Context, bid entity.Bid) (entity.Bid, error) {
	err := bw.dbConn.QueryRow(ctx,
		`INSERT INTO bid(product_id, user_id, amount, status, bid_time) VALUES($1, $2, $3, $4, $5)
		RETURNING id`,
		bid.ProductID, bid.UserID, bid.Amount, bid.Status, bid.BidTime).Scan(&bid.ID)
	if err != nil {
		return bid, stacktrace.Propagate(err, "[bidWriter][Insert][QueryRow]")
	}

	return bid, nil
}

func (bw *bidWriter) UpdateAmount(ctx context.Context, bidID, newAmount int) error {
	_, err := bw.dbConn.Exec(ctx,
		`UPDATE bid SET amount = $1 WHERE id = $2`,
		newAmount, bidID)
	if err != nil {
		return stacktrace.Propagate(err, "[bidWriter][UpdateAmount][Exec]")
	}
	return nil
}

func (bw *bidWriter) UpdateStatus(ctx context.Context, bidID, status int) error {
	_, err := bw.dbConn.Exec(ctx,
		`UPDATE bid SET status = $1 WHERE id = $2`,
		status, bidID)
	if err != nil {
		return stacktrace.Propagate(err, "[bidWriter][UpdateStatus][Exec]")
	}
	return nil
}

func (bw *bidWriter) WithTx(tx db.Tx) src.BidWriterRepo {
	return NewBidWriter(tx)
}
