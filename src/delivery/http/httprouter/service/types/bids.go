package types

import (
	"github.com/idzharbae/quickbid/src/entity"
)

type ListUserBidsResponse struct {
	Bids []BidWithProduct `json:"bids"`
}

type ListProductBidsResponse struct {
	Bids []BidWithBidder `json:"bids"`
}

type BidWithProduct struct {
	Amount  int        `json:"amount"`
	Status  string     `json:"status"`
	Product BidProduct `json:"product"`
}

type BidProduct struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type BidWithBidder struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	ActiveBid bool   `json:"active_bid"`
}

func BidWithProductFromEntity(bid entity.BidWithProduct) BidWithProduct {
	return BidWithProduct{
		Amount: bid.Amount,
		Status: bidsStatusToString(bid.Status),
		Product: BidProduct{
			Name:     bid.Product.Name,
			ImageURL: bid.Product.ImageURL,
		},
	}
}

func ListBidResponseFromBidsWithProduct(bids []entity.BidWithProduct) ListUserBidsResponse {
	var res ListUserBidsResponse
	for _, bid := range bids {
		res.Bids = append(res.Bids, BidWithProductFromEntity(bid))
	}
	return res
}

func BidWithBidderFromEntity(bid entity.BidWithBidder) BidWithBidder {
	return BidWithBidder{
		Amount:    bid.Amount,
		Name:      bid.Bidder.Name,
		ID:        bid.ID,
		ActiveBid: bid.Status == entity.BidStatusActive,
	}
}

func ListProductBidsResponseFromEntity(bids []entity.BidWithBidder) ListProductBidsResponse {
	var res ListProductBidsResponse
	for _, bid := range bids {
		res.Bids = append(res.Bids, BidWithBidderFromEntity(bid))
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
