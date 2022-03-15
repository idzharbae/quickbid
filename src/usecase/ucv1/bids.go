package ucv1

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"

	"github.com/palantir/stacktrace"
)

type bidUC struct {
	bidReader src.BidReaderRepo
}

func NewBidUC(bidReader src.BidReaderRepo) src.BidUC {
	return &bidUC{bidReader: bidReader}
}

func (b *bidUC) ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.BidWithProduct, error) {
	bids, err := b.bidReader.ListUserBiddedProducts(ctx, req.UserID, req.Page, req.Limit)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidUC][ListUserBiddedProducts]")
	}

	return bids, nil
}
