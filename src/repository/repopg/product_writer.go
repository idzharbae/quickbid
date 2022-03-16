package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type productWriter struct {
	dbConn db.Connection
}

func NewProductWriter(dbConn db.Connection) src.ProductWriterRepo {
	return &productWriter{dbConn: dbConn}
}

func (at *productWriter) InsertProduct(ctx context.Context, product entity.Product) error {
	_, err := at.dbConn.Exec(ctx, InsertProductQuery, product.Name, product.InitialPrice, product.StartBidDate, product.EndBidDate, product.OwnerUserID, product.LastBidID, product.ImageURL, product.Status, product.BidIncrement)
	if err != nil {
		return err
	}
	return nil
}

func (at *productWriter) UpdateLastBidID(ctx context.Context, productID, lastBidID int) error {
	_, err := at.dbConn.Exec(ctx, "UPDATE products SET last_bid_id = $1 WHERE id = $2", lastBidID, productID)
	if err != nil {
		return stacktrace.Propagate(err, "[productWriter][UpdateLastBidID][dbConn][Exec]")
	}

	return nil
}

func (at *productWriter) WithTx(tx db.Tx) src.ProductWriterRepo {
	return NewProductWriter(tx)
}
