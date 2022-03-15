package ucv1

import (
	"context"
	"fmt"
	"time"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/bridge/transactioner"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/jackc/pgx/v4"
)

type attendanceUC struct {
	attendanceWriter src.AttendanceWriterRepo
	attendanceReader src.AttendanceReaderRepo

	txner transactioner.Transactioner
}

func NewAttendanceV1(attendanceWriter src.AttendanceWriterRepo, attendanceReader src.AttendanceReaderRepo, txner transactioner.Transactioner) *attendanceUC {
	return &attendanceUC{
		attendanceWriter: attendanceWriter,
		attendanceReader: attendanceReader,
		txner:            txner,
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

func (ucv1 *attendanceUC) AttendBulk(ctx context.Context, names []string) error {
	err := ucv1.txner.DoWithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		reader := ucv1.attendanceReader.WithTx(tx)
		writer := ucv1.attendanceWriter.WithTx(tx)

		for _, name := range names {
			userAttendance, err := reader.GetByName(ctx, name)
			if err != nil && err != pgx.ErrNoRows {
				return err
			}

			if err == nil {
				return fmt.Errorf("[AttendBulk] user %s already attended at %v", name, userAttendance.AttendanceTime)
			}

			err = writer.Insert(ctx, entity.Attendance{
				Name:           name,
				AttendanceTime: time.Now(),
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
