package src

import "context"

//go:generate mockgen -destination=usecase/ucmock/usecase_mock.go -package=ucmock -source=usecase.go

type AttendanceUC interface {
	Attend(ctx context.Context, name string) error
	AttendBulk(ctx context.Context, names []string) error
}
