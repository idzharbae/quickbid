package ucv1

import (
	"context"
	"fmt"
	"time"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/jackc/pgx/v4"
)

type attendanceUC struct {
	attendanceWriter src.AttendanceWriterRepo
	attendanceReader src.AttendanceReaderRepo
}

func NewAttendanceV1(attendanceWriter src.AttendanceWriterRepo, attendanceReader src.AttendanceReaderRepo) *attendanceUC {
	return &attendanceUC{
		attendanceWriter: attendanceWriter,
		attendanceReader: attendanceReader,
	}
}

func (ucv1 *attendanceUC) Attend(ctx context.Context, name string) error {
	userAttendance, err := ucv1.attendanceReader.GetByName(ctx, name)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if err == nil {
		return fmt.Errorf("user %s already attended at %v", name, userAttendance.AttendanceTime)
	}

	err = ucv1.attendanceWriter.Insert(ctx, entity.Attendance{
		Name:           name,
		AttendanceTime: time.Now(),
	})

	return err
}
