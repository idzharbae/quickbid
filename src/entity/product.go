package entity

import "time"

const (
	ProductStatusActive   = 1
	ProductStatusInactive = 2
)

var (
	statusStringMap = map[int]string{
		ProductStatusActive:   "active",
		ProductStatusInactive: "inactive",
	}
)

type Product struct {
	ID           int
	Name         string
	InitialPrice int
	StartBidDate time.Time
	EndBidDate   time.Time
	OwnerUserID  int
	LastBidID    int
	ImageURL     string
	Status       int
	BidIncrement int
}

type ProductWithSeller struct {
	Product
	Seller User
}
