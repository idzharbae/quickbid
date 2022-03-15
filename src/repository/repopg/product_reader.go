package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type productReader struct {
	dbConn db.Connection
}

func NewProductReader(dbConn db.Connection) src.ProductReaderRepo {
	return &productReader{dbConn: dbConn}
}

func (p *productReader) GetByIDWithSeller(ctx context.Context, productID int) (entity.ProductWithSeller, error) {
	var res entity.ProductWithSeller

	err := p.dbConn.QueryRow(ctx,
		`SELECT p.id, p.name, p.initial_price, p.start_bid_date, p.end_bid_date, p.owner_user_id, COALESCE(p.last_bid_id, 0), p.image_url, p.status, p.bid_increment,
			u.id, u.email, u.name
			FROM products p JOIN users u ON p.owner_user_id = u.id
			WHERE p.id = $1`,
		productID,
	).Scan(&res.ID, &res.Name, &res.InitialPrice, &res.StartBidDate, &res.EndBidDate, &res.OwnerUserID, &res.LastBidID, &res.ImageURL, &res.Status, &res.BidIncrement,
		&res.Seller.ID, &res.Seller.Email, &res.Seller.Name)
	if err != nil {
		return res, stacktrace.Propagate(err, "[ListUserBiddedProducts][bidReader]")
	}

	return res, nil
}

func (p *productReader) WithTx(tx db.Tx) src.ProductReaderRepo {
	return NewProductReader(tx)
}
