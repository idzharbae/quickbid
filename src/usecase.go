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
	ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.BidWithProduct, error)
}

type ProductUC interface {
	GetByIDWithSeller(ctx context.Context, productID int) (entity.ProductWithSeller, error)
	UploadProduct(ctx context.Context, req requests.UploadProductRequest) error
	GetProductByOwnerUserID(ctx context.Context, req requests.GetProductByOwnerUserIDRequest) ([]entity.Product, error)
}
