package entity

import "time"

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
