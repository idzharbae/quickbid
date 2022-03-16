package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type productReader struct {
	dbConn  db.Connection
	isInTrx bool
}

func NewProductReader(dbConn db.Connection) src.ProductReaderRepo {
	return &productReader{dbConn: dbConn}
}

func newProductReaderTx(dbConn db.Connection) src.ProductReaderRepo {
	return &productReader{dbConn: dbConn, isInTrx: true}
}

func (p *productReader) GetByIDAndLock(ctx context.Context, productID int) (entity.Product, error) {
	var res entity.Product

	if !p.isInTrx {
		return res, stacktrace.NewError("can't lock outside transaction")
	}

	err := p.dbConn.QueryRow(ctx,
		`SELECT p.id, p.name, p.initial_price, p.start_bid_date, p.end_bid_date, p.owner_user_id, COALESCE(p.last_bid_id, 0), p.image_url, p.status, p.bid_increment
			FROM products p
			WHERE p.id = $1
			FOR UPDATE`,
		productID,
	).Scan(&res.ID, &res.Name, &res.InitialPrice, &res.StartBidDate, &res.EndBidDate, &res.OwnerUserID, &res.LastBidID, &res.ImageURL, &res.Status, &res.BidIncrement)
	if err != nil {
		return res, stacktrace.Propagate(err, "[productReader][GetByIDAndLock]")
	}

	return res, nil
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
		return res, stacktrace.Propagate(err, "[productReader][GetByIDWithSeller]")
	}

	return res, nil
}

func (at *productReader) GetProductByOwnerUserID(ctx context.Context, ownerUserId int) ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := at.dbConn.Query(ctx, GetProductByOwnerUserIDQuery, ownerUserId)
	if err != nil {
		return []entity.Product{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.InitialPrice,
			&product.StartBidDate,
			&product.EndBidDate,
			&product.OwnerUserID,
			&product.LastBidID,
			&product.ImageURL,
			&product.Status,
			&product.BidIncrement,
		); err != nil {
			return []entity.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *productReader) WithTx(tx db.Tx) src.ProductReaderRepo {
	return newProductReaderTx(tx)
}
