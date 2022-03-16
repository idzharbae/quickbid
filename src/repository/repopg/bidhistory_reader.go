package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type bidHistoryReader struct {
	dbConn db.Connection
}

func NewBidHistoryReader(dbConn db.Connection) src.BidHistoryReaderRepo {
	return &bidHistoryReader{dbConn: dbConn}
}

func (bhr *bidHistoryReader) GetByID(ctx context.Context, id int) (entity.Bid, error) {
	var res entity.Bid
	err := bhr.dbConn.QueryRow(ctx, `SELECT id, user_id, amount, product_id, bid_time
			FROM bid_history WHERE id = $1`, id).
		Scan(&res.ID, &res.UserID, &res.Amount, &res.ProductID, &res.BidTime)
	if err != nil {
		return res, stacktrace.Propagate(err, "[bidHistoryReader][GetByID]")
	}

	return res, nil
}

func (bhr *bidHistoryReader) ListByProductID(ctx context.Context, productID int, page int, limit int) ([]entity.BidWithBidder, error) {
	offset := page * limit

	rows, err := bhr.dbConn.Query(ctx,
		`SELECT b.id, b.user_id, b.amount, COALESCE(b.bid_time, '0001-01-01T00:00:00Z'::timestamp), b.product_id,
			u.id, u.name, u.email
			FROM bid b JOIN users u ON b.user_id = u.id 
			WHERE b.product_id = $1 
			ORDER BY b.id desc LIMIT $2 OFFSET $3`,
		productID, limit, offset,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidHistoryReader][ListByProductID]")
	}

	defer rows.Close()

	result := make([]entity.BidWithBidder, 0, rows.RowsAffected())

	for rows.Next() {
		var bid entity.BidWithBidder
		err := rows.Scan(&bid.ID, &bid.UserID, &bid.Amount, &bid.BidTime, &bid.ProductID,
			&bid.Bidder.ID, &bid.Bidder.Name, &bid.Bidder.Email)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[bidHistoryReader][ListByProductID]")
		}

		result = append(result, bid)
	}

	return result, nil
}

func (bhr *bidHistoryReader) WithTx(tx db.Tx) src.BidHistoryReaderRepo {
	return NewBidHistoryReader(tx)
}
