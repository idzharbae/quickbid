package app

import (
	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/usecase/ucv1"
)

type UseCases struct {
	AttendanceUC src.AttendanceUC
	BidUC        src.BidUC
}

func newUseCases(repos *Repositories, bridges *Bridges) *UseCases {
	return &UseCases{
		AttendanceUC: ucv1.NewAttendanceV1(repos.AttendanceWriter, repos.AttendanceReader, bridges.transactioner),
		BidUC:        ucv1.NewBidUC(repos.BidReader),
	}
}
