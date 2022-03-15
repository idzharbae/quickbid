package types

import "github.com/idzharbae/quickbid/src/entity"

type ListBidResponse struct {
	Bids []Bid `json:"bids"`
}

type Bid struct {
	Amount  int        `json:"amount"`
	Status  string     `json:"status"`
	Product BidProduct `json:"product"`
}

type BidProduct struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func BidFromEntity(bid entity.BidWithProduct) Bid {
	return Bid{
		Amount: bid.Amount,
		Status: bidsStatusToString(bid.Status),
		Product: BidProduct{
			Name:     bid.Product.Name,
			ImageURL: bid.Product.ImageURL,
		},
	}
}

func ListBidResponseFromEntity(bids []entity.BidWithProduct) ListBidResponse {
	var res ListBidResponse
	for _, bid := range bids {
		res.Bids = append(res.Bids, BidFromEntity(bid))
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
