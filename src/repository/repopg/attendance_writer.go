package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
)

type attendanceWriter struct {
	dbConn db.Connection
}

func NewAttendanceWriter(dbConn db.Connection) src.AttendanceWriterRepo {
	return &attendanceWriter{dbConn: dbConn}
}

func (at *attendanceWriter) Insert(ctx context.Context, req entity.Attendance) error {
	_, err := at.dbConn.Exec(ctx, "INSERT INTO attendance(name, attendance_time) VALUES($1, $2)", req.Name, req.AttendanceTime)
	return err
}

func (at *attendanceWriter) WithTx(tx db.Tx) src.AttendanceWriterRepo {
	return NewAttendanceWriter(tx)
}
