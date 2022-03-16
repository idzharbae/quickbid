package ucv1

import (
	"context"

	"github.com/ggwhite/go-masker"
	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"

	"github.com/palantir/stacktrace"
)

type bidUC struct {
	bidReader        src.BidReaderRepo
	bidHistoryReader src.BidHistoryReaderRepo
}

func NewBidUC(bidReader src.BidReaderRepo, bidHistoryReader src.BidHistoryReaderRepo) src.BidUC {
	return &bidUC{
		bidReader:        bidReader,
		bidHistoryReader: bidHistoryReader,
	}
}

func (b *bidUC) ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.BidWithProduct, error) {
	bids, err := b.bidReader.ListUserBiddedProducts(ctx, req.UserID, req.Page, req.Limit)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidUC][ListUserBiddedProducts]")
	}

	return bids, nil
}

func (b *bidUC) ListByProduct(ctx context.Context, req requests.ListBidsByProductRequest) ([]entity.BidWithBidder, error) {
	bids, err := b.bidHistoryReader.ListByProductID(ctx, req.ProductID, req.Page, req.Limit)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidUC][ListByProduct]")
	}

	// Mask bidder name
	for i := range bids {
		bids[i].Bidder.Name = masker.Name(bids[i].Bidder.Name)
	}

	return bids, nil
}
