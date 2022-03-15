package entity

import "time"

type Bid struct {
	ID     int
	UserID int
	Amount int
	Status int
}

type BidHistory struct {
	ID      int
	UserID  int
	Amount  int
	BidTime time.Time
}
