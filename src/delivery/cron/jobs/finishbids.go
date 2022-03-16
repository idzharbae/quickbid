package jobs

import (
	"context"
	"log"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"
)

type BidFinisher struct {
	bidUC     src.BidUC
	productUC src.ProductUC
}

func NewBidFinisher(bidUC src.BidUC, productUC src.ProductUC) *BidFinisher {
	return &BidFinisher{
		bidUC:     bidUC,
		productUC: productUC,
	}
}

func (bf *BidFinisher) Run() {
	ctx := context.Background()
	finishedProducts, err := bf.productUC.GetFinishedProducts(ctx, 0, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range finishedProducts {
		wonBidId := product.LastBidID
		wonBid, err := bf.bidUC.GetByID(ctx, wonBidId)
		if err != nil {
			log.Fatal(err)
		}

		if wonBid.Status != entity.BidStatusWon && wonBid.Status != entity.BidStatusRefunded {
			err = bf.bidUC.SetAsWinner(ctx, wonBidId)
			if err != nil {
				log.Fatal(err)
			}
		}

		page := 0
		var lostBids []entity.Bid

		for page == 0 || len(lostBids) == 0 {
			lostBids, err = bf.bidUC.ListByProductAndStatus(ctx, requests.ListBidsByProductAndStatusRequest{
				ProductID: product.ID,
				Page:      page,
				Limit:     10,
				Status:    entity.BidStatusInactive,
			})
			if err != nil {
				log.Fatal(err)
			}

			if len(lostBids) == 0 {
				break
			}

			var lostBidIds []int

			for _, bid := range lostBids {
				lostBidIds = append(lostBidIds, bid.ID)
			}

			err = bf.bidUC.SetAsLoserBulk(ctx, lostBidIds)
			if err != nil {
				log.Fatal(err)
			}

			page += 1
		}

		bf.productUC.SetAsFinished(ctx, product.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
