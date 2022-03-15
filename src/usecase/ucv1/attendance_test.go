package ucv1_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/quickbid/src/bridge/transactioner/txnermock"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/repository/repomock"
	"github.com/idzharbae/quickbid/src/usecase/ucv1"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestUCV1_Attend(t *testing.T) {
	type args struct {
		Name string
	}

	testCases := []struct {
		args    args
		desc    string
		wantErr error
		mock    func(t *testing.T, attendanceWriterMock *repomock.MockAttendanceWriterRepo, attendanceReaderMock *repomock.MockAttendanceReaderRepo)
	}{
		{
			desc:    "Reader returns error (not err no rows), must return the error",
			args:    args{Name: "test"},
			wantErr: errors.New("error"),
			mock: func(t *testing.T, attendanceWriterMock *repomock.MockAttendanceWriterRepo, attendanceReaderMock *repomock.MockAttendanceReaderRepo) {
				gomock.InOrder(
					attendanceReaderMock.EXPECT().GetByName(gomock.Any(), "test").Return(entity.Attendance{}, errors.New("error")),
				)
			},
		},
		{
			desc:    "Reader returns no rows error, then writer returns error, must call writer with correct param and return the error",
			args:    args{Name: "test"},
			wantErr: errors.New("error"),
			mock: func(t *testing.T, attendanceWriterMock *repomock.MockAttendanceWriterRepo, attendanceReaderMock *repomock.MockAttendanceReaderRepo) {
				gomock.InOrder(
					attendanceReaderMock.EXPECT().GetByName(gomock.Any(), "test").Return(entity.Attendance{}, pgx.ErrNoRows),
					attendanceWriterMock.EXPECT().Insert(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req entity.Attendance) error {
						assert.Equal(t, "test", req.Name)
						return errors.New("error")
					}),
				)
			},
		},
		{
			desc:    "Reader returns no error, must return error because user already attended",
			args:    args{Name: "test"},
			wantErr: fmt.Errorf("user %s already attended at %v", "test", time.Time{}),
			mock: func(t *testing.T, attendanceWriterMock *repomock.MockAttendanceWriterRepo, attendanceReaderMock *repomock.MockAttendanceReaderRepo) {
				gomock.InOrder(
					attendanceReaderMock.EXPECT().GetByName(gomock.Any(), "test").Return(entity.Attendance{Name: "test"}, nil),
				)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			attendanceReaderMock := repomock.NewMockAttendanceReaderRepo(ctrl)
			attendanceWriterMock := repomock.NewMockAttendanceWriterRepo(ctrl)
			txnerMock := txnermock.NewMockTransactioner(ctrl)

			tC.mock(t, attendanceWriterMock, attendanceReaderMock)

			uc := ucv1.NewAttendanceV1(attendanceWriterMock, attendanceReaderMock, txnerMock)
			err := uc.Attend(context.Background(), tC.args.Name)
			if !reflect.DeepEqual(err, tC.wantErr) {
				t.Fatalf("[ucv1][Attend] Want err: %v, got err: %v", tC.wantErr, err)
			}
		})
	}
}
