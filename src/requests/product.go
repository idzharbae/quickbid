package requests

import (
	"time"
)

type (
	UploadProductRequest struct {
		Name         string    `json:"name"`
		InitialPrice int       `json:"initial_price"`
		StartBidDate time.Time `json:"start_time"`
		EndBidDate   time.Time `json:"end_time"`
		OwnerUserID  int       `json:"seller_id"`
		ImageURL     string    `json:"image_url"`
		BidIncrement int       `json:"bid_increment"`
	}

	GetProductByOwnerUserIDRequest struct {
		OwnerUserID int `json:"owner_user_id"`
	}
)
