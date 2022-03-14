package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/repository/repopg/pgx"
)

type attendanceWriter struct {
	pgx pgx.Pgx
}

func NewAttendanceWriter(pgDriver pgx.Pgx) *attendanceWriter {
	return &attendanceWriter{pgx: pgDriver}
}

func (at *attendanceWriter) Insert(ctx context.Context, req entity.Attendance) error {
	_, err := at.pgx.Exec(ctx, "INSERT INTO attendance(name, attendance_time) VALUES($1, $2)", req.Name, req.AttendanceTime)
	return err
}
