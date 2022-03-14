package src

import (
	"context"

	"github.com/idzharbae/quickbid/src/entity"
)

//go:generate mockgen -destination=repository/repomock/repo_mock.go -package=repomock -source=repository.go

type AttendanceReaderRepo interface {
	GetByName(ctx context.Context, name string) (entity.Attendance, error)
}

type AttendanceWriterRepo interface {
	Insert(ctx context.Context, attendance entity.Attendance) error
}
