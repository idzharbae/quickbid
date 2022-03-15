package entity

import "time"

const (
	BidStatusInactive = iota
	BidStatusActive
	BidStatusWon
	BidStatusRefunded
)

type Bid struct {
	ID     int
	UserID int
	Amount int
	Status int
}

type BidWithProduct struct {
	Bid
	Product Product
}

type BidHistory struct {
	ID      int
	UserID  int
	Amount  int
	BidTime time.Time
}
