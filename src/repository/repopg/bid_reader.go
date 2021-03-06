package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/lib/pq"
	"github.com/palantir/stacktrace"
)

type bidReader struct {
	dbConn db.Connection
}

func NewBidReader(dbConn db.Connection) src.BidReaderRepo {
	return &bidReader{dbConn: dbConn}
}

func (bhr *bidReader) GetByID(ctx context.Context, id int) (entity.Bid, error) {
	var res entity.Bid
	err := bhr.dbConn.QueryRow(ctx, `SELECT id, user_id, amount, product_id, bid_time, status
			FROM bid WHERE id = $1`, id).
		Scan(&res.ID, &res.UserID, &res.Amount, &res.ProductID, &res.BidTime, &res.Status)
	if err != nil {
		return res, stacktrace.Propagate(err, "[bidReader][GetByID]")
	}

	return res, nil
}

func (br *bidReader) ListByProductIDAndStatus(ctx context.Context, productID, status, limit, page int) ([]entity.Bid, error) {
	rows, err := br.dbConn.Query(ctx,
		`SELECT b.id, b.user_id, b.amount, b.status, COALESCE(b.bid_time, '0001-01-01T00:00:00Z'::timestamp), b.product_id
			FROM bid b JOIN products p ON b.product_id = p.id 
			WHERE b.product_id = $1 AND b.status = $2 
			LIMIT $3 OFFSET $4`,
		productID, status, limit, page*limit,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidReader][ListByProductAndStatus]")
	}

	defer rows.Close()

	result := make([]entity.Bid, 0, rows.RowsAffected())

	for rows.Next() {
		var bid entity.Bid
		err := rows.Scan(&bid.ID, &bid.UserID, &bid.Amount, &bid.Status, &bid.BidTime, &bid.ProductID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[bidReader][ListByProductAndStatus]")
		}

		result = append(result, bid)
	}

	return result, nil
}

func (bhr *bidReader) GetByIDWithProduct(ctx context.Context, id int) (entity.BidWithProduct, error) {
	var bid entity.BidWithProduct
	err := bhr.dbConn.QueryRow(ctx, `SELECT b.id, b.user_id, b.amount, b.status, COALESCE(b.bid_time, '0001-01-01T00:00:00Z'::timestamp), b.product_id,
			p.id, p.name, p.image_url, p.initial_price, p.bid_increment, p.owner_user_id, p.start_bid_date, p.end_bid_date
			FROM bid b JOIN products p ON b.product_id = p.id 
			WHERE b.id = $1`, id).
		Scan(&bid.ID, &bid.UserID, &bid.Amount, &bid.Status, &bid.BidTime, &bid.ProductID,
			&bid.Product.ID, &bid.Product.Name, &bid.Product.ImageURL, &bid.Product.InitialPrice, &bid.Product.BidIncrement, &bid.Product.OwnerUserID, &bid.Product.StartBidDate, &bid.Product.EndBidDate)
	if err != nil {
		return bid, stacktrace.Propagate(err, "[bidReader][GetByIDWithProduct]")
	}

	return bid, nil
}

func (br *bidReader) GetByIDsWithProduct(ctx context.Context, ids []int) ([]entity.BidWithProduct, error) {
	rows, err := br.dbConn.Query(ctx,
		`SELECT b.id, b.user_id, b.amount, b.status, COALESCE(b.bid_time, '0001-01-01T00:00:00Z'::timestamp), b.product_id,
			p.id, p.name, p.image_url, p.initial_price, p.bid_increment, p.owner_user_id, p.start_bid_date, p.end_bid_date
			FROM bid b JOIN products p ON b.product_id = p.id 
			WHERE b.id = any($1)`,
		pq.Array(ids),
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidReader][ListUserBiddedProducts]")
	}

	defer rows.Close()

	result := make([]entity.BidWithProduct, 0, rows.RowsAffected())

	for rows.Next() {
		var bid entity.BidWithProduct
		err := rows.Scan(&bid.ID, &bid.UserID, &bid.Amount, &bid.Status, &bid.BidTime, &bid.ProductID,
			&bid.Product.ID, &bid.Product.Name, &bid.Product.ImageURL, &bid.Product.InitialPrice, &bid.Product.BidIncrement, &bid.Product.OwnerUserID, &bid.Product.StartBidDate, &bid.Product.EndBidDate)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[bidReader][ListUserBiddedProducts]")
		}

		result = append(result, bid)
	}

	return result, nil
}

func (br *bidReader) ListUserBiddedProducts(ctx context.Context, userID int, page int, limit int) ([]entity.BidWithProduct, error) {
	offset := page * limit

	rows, err := br.dbConn.Query(ctx,
		`SELECT b.id, b.user_id, b.amount, b.status, COALESCE(b.bid_time, '0001-01-01T00:00:00Z'::timestamp), b.product_id,
			p.id, p.name, p.image_url, p.initial_price, p.bid_increment, p.owner_user_id, p.start_bid_date, p.end_bid_date
			FROM bid b JOIN products p ON b.product_id = p.id 
			WHERE b.user_id = $1 
			ORDER BY b.id desc LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidReader][ListUserBiddedProducts]")
	}

	defer rows.Close()

	result := make([]entity.BidWithProduct, 0, rows.RowsAffected())

	for rows.Next() {
		var bid entity.BidWithProduct
		err := rows.Scan(&bid.ID, &bid.UserID, &bid.Amount, &bid.Status, &bid.BidTime, &bid.ProductID,
			&bid.Product.ID, &bid.Product.Name, &bid.Product.ImageURL, &bid.Product.InitialPrice, &bid.Product.BidIncrement, &bid.Product.OwnerUserID, &bid.Product.StartBidDate, &bid.Product.EndBidDate)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[bidReader][ListUserBiddedProducts]")
		}

		result = append(result, bid)
	}

	return result, nil
}

func (br *bidReader) GetByUserIDAndProductID(ctx context.Context, userID, productID int) (entity.Bid, error) {
	var res entity.Bid
	err := br.dbConn.QueryRow(ctx, `SELECT id, user_id, amount, product_id, status, COALESCE(bid_time, '0001-01-01T00:00:00Z'::timestamp)
			FROM bid WHERE user_id = $1 AND product_id = $2`, userID, productID).
		Scan(&res.ID, &res.UserID, &res.Amount, &res.ProductID, &res.Status, &res.BidTime)
	if err != nil {
		return res, stacktrace.Propagate(err, "[bidReader][GetByUserIDAndProductID]")
	}

	return res, nil
}

func (br *bidReader) WithTx(tx db.Tx) src.BidReaderRepo {
	return NewBidReader(tx)
}
