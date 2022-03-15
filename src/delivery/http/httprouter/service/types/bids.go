package types

import "github.com/idzharbae/quickbid/src/entity"

type Bid struct {
	Amount  int     `json:"amount"`
	Status  string  `json:"status"`
	Product Product `json:"product"`
}

func BidResponseFromEntity(bid entity.Bid) Bid {
	return Bid{
		Amount: bid.Amount,
		Status: bidsStatusToString(bid.Status),
		Product: Product{
			Name:     bid.Product.Name,
			ImageURL: bid.Product.ImageURL,
		},
	}
}

func BidsRepsonseFromEntity(bids []entity.Bid) []Bid {
	var res []Bid
	for _, bid := range bids {
		res = append(res, BidResponseFromEntity(bid))
	}
	return res
}

func bidsStatusToString(bidStatus int) string {
	switch bidStatus {
	case entity.BidStatusActive:
		return "active"
	case entity.BidStatusInactive:
		return "inactive"
	case entity.BidStatusRefunded:
		return "refunded"
	case entity.BidStatusWon:
		return "won"
	default:
		return ""
	}
}
