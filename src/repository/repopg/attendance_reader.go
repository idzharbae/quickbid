package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
)

type attendanceReader struct {
	dbConn db.Connection
}

func NewAttendanceReader(dbConn db.Connection) src.AttendanceReaderRepo {
	return &attendanceReader{dbConn: dbConn}
}

func (at *attendanceReader) GetByName(ctx context.Context, name string) (res entity.Attendance, err error) {
	err = at.dbConn.
		QueryRow(ctx, "SELECT name, attendance_time FROM attendance WHERE name = $1", name).
		Scan(&res.Name, &res.AttendanceTime)
	if err != nil {
		return entity.Attendance{}, err
	}

	return res, nil
}

func (at *attendanceReader) WithTx(tx db.Tx) src.AttendanceReaderRepo {
	return NewAttendanceReader(tx)
}
