package src

import (
	"context"

	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"
)

//go:generate mockgen -destination=usecase/ucmock/usecase_mock.go -package=ucmock -source=usecase.go

type AttendanceUC interface {
	Attend(ctx context.Context, name string) error
	AttendBulk(ctx context.Context, names []string) error
}

type BidUC interface {
	ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.Bid, error)
}
