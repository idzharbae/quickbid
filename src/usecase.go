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
	GetByID(ctx context.Context, id int) (entity.Bid, error)
	ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.BidWithProduct, error)
	ListByProduct(ctx context.Context, req requests.ListBidsByProductRequest) ([]entity.BidWithBidder, error)
	ListByProductAndStatus(ctx context.Context, req requests.ListBidsByProductAndStatusRequest) ([]entity.Bid, error)
	SetAsWinner(ctx context.Context, bidID int) error
	SetAsLoserBulk(ctx context.Context, bidIDs []int) error

	BidProduct(ctx context.Context, req requests.BidProductRequest) (entity.Bid, error)
}

type ProductUC interface {
	GetByIDWithSeller(ctx context.Context, productID int) (entity.ProductWithSeller, error)
	UploadProduct(ctx context.Context, req requests.UploadProductRequest) error
	GetProductByOwnerUserID(ctx context.Context, req requests.GetProductByOwnerUserIDRequest) ([]entity.Product, error)
	GetFinishedProducts(ctx context.Context, page, limit int) ([]entity.Product, error)
	SetAsFinished(ctx context.Context, productID int) error
}
