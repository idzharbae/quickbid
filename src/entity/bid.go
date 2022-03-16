package entity

import "time"

const (
	BidStatusInactive = iota
	BidStatusActive
	BidStatusWon
	BidStatusRefunded
)

type Bid struct {
	ID      int
	UserID  int
	Amount  int
	Status  int
	BidTime time.Time
}

type BidWithProduct struct {
	Bid
	Product Product
}

type BidWithBidder struct {
	Bid
	Bidder User
}

type BidHistory struct {
	ID      int
	UserID  int
	Amount  int
	BidTime time.Time
}
