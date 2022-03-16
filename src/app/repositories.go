package app

import (
	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db/pgx"
	"github.com/idzharbae/quickbid/src/repository/repopg"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repositories struct {
	AttendanceWriter src.AttendanceWriterRepo
	AttendanceReader src.AttendanceReaderRepo
	BidReader        src.BidReaderRepo
	BidWriter        src.BidWriterRepo
	BidHistoryReader src.BidHistoryReaderRepo
	BidHistoryWriter src.BidHistoryWriterRepo
	ProductReader    src.ProductReaderRepo
	ProductWriter    src.ProductWriterRepo
	WalletReader     src.WalletReaderRepo
	WalletWriter     src.WalletWriterRepo
}

func newRepositories(pgxPool *pgxpool.Pool) *Repositories {
	pgxDriver := pgx.NewPgxDriver(pgxPool)

	return &Repositories{
		AttendanceWriter: repopg.NewAttendanceWriter(pgxDriver),
		AttendanceReader: repopg.NewAttendanceReader(pgxDriver),
		BidReader:        repopg.NewBidReader(pgxDriver),
		BidWriter:        repopg.NewBidWriter(pgxDriver),
		BidHistoryReader: repopg.NewBidHistoryReader(pgxDriver),
		BidHistoryWriter: repopg.NewBidHistoryWriter(pgxDriver),
		ProductReader:    repopg.NewProductReader(pgxDriver),
		ProductWriter:    repopg.NewProductWriter(pgxDriver),
		WalletReader:     repopg.NewWalletReader(pgxDriver),
		WalletWriter:     repopg.NewWalletWriter(pgxDriver),
	}
}
