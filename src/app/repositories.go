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
	ProductReader    src.ProductReaderRepo
}

func newRepositories(pgxPool *pgxpool.Pool) *Repositories {
	pgxDriver := pgx.NewPgxDriver(pgxPool)

	return &Repositories{
		AttendanceWriter: repopg.NewAttendanceWriter(pgxDriver),
		AttendanceReader: repopg.NewAttendanceReader(pgxDriver),
		BidReader:        repopg.NewBidReader(pgxDriver),
		ProductReader:    repopg.NewProductReader(pgxDriver),
	}
}
