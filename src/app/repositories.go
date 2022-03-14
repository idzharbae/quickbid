package app

import (
	"fmt"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/app/config"
	"github.com/idzharbae/quickbid/src/repository/repopg"
	"github.com/idzharbae/quickbid/src/repository/repopg/pgx/pgxdriver"
)

type Repositories struct {
	AttendanceWriter src.AttendanceWriterRepo
	AttendanceReader src.AttendanceReaderRepo
}

func newRepositories(cfg config.Config) *Repositories {
	pgDriver := pgxdriver.NewPgxDriver(fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Address, cfg.DB.Port, cfg.DB.DBName))

	return &Repositories{
		AttendanceWriter: repopg.NewAttendanceWriter(pgDriver),
		AttendanceReader: repopg.NewAttendanceReader(pgDriver),
	}
}
