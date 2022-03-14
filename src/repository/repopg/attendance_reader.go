package repopg

import (
	"context"

	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/repository/repopg/pgx"
)

type attendanceReader struct {
	pgx pgx.Pgx
}

func NewAttendanceReader(pgDriver pgx.Pgx) *attendanceReader {
	return &attendanceReader{pgx: pgDriver}
}

func (at *attendanceReader) GetByName(ctx context.Context, name string) (res entity.Attendance, err error) {
	row, err := at.pgx.QueryRow(ctx, "SELECT name, attendance_time FROM attendance WHERE name = $1", name)
	if err != nil {
		return entity.Attendance{}, err
	}

	err = row.Scan(&res.Name, &res.AttendanceTime)
	if err != nil {
		return entity.Attendance{}, err
	}

	return entity.Attendance{}, nil
}
