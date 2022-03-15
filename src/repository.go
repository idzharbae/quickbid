package src

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
)

//go:generate mockgen -destination=repository/repomock/repo_mock.go -package=repomock -source=repository.go

type AttendanceReaderRepo interface {
	GetByName(ctx context.Context, name string) (entity.Attendance, error)

	WithTx(db.Tx) AttendanceReaderRepo
}

type AttendanceWriterRepo interface {
	Insert(ctx context.Context, attendance entity.Attendance) error

	WithTx(db.Tx) AttendanceWriterRepo
}

type BidReaderRepo interface {
	ListUserBiddedProducts(ctx context.Context, userID int, page int, limit int) ([]entity.Bid, error)

	WithTx(db.Tx) BidReaderRepo
}
