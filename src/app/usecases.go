package app

import (
	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/usecase/ucv1"
)

type UseCases struct {
	AttendanceUC src.AttendanceUC
}

func newUseCases(repos *Repositories) *UseCases {
	return &UseCases{
		AttendanceUC: ucv1.NewAttendanceV1(repos.AttendanceWriter, repos.AttendanceReader),
	}
}
