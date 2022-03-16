package types

import (
	"time"

	"github.com/idzharbae/quickbid/src/entity"
)

type GetProductResponse struct {
	Name         string    `json:"name"`
	InitialPrice int       `json:"initial_price"`
	BidIncrement int       `json:"bid_increment"`
	ImageURL     string    `json:"image_url"`
	SellerName   string    `json:"seller_name"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}

func GetProductResponseFromProductAndSellerEntity(product entity.ProductWithSeller) GetProductResponse {
	return GetProductResponse{
		Name:         product.Name,
		InitialPrice: product.InitialPrice,
		BidIncrement: product.BidIncrement,
		ImageURL:     product.ImageURL,
		SellerName:   product.Seller.Name,
		StartTime:    product.StartBidDate,
		EndTime:      product.EndBidDate,
	}
}

type Product struct {
	Name         string    `json:"name"`
	InitialPrice int       `json:"initial_price"`
	StartBidDate time.Time `json:"start_time"`
	EndBidDate   time.Time `json:"end_time"`
	OwnerUserID  int       `json:"seller_id"`
	ImageURL     string    `json:"image_url"`
	BidIncrement int       `json:"bid_increment"`
}
