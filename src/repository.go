package src

import (
	"context"

	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/entity"
)

//go:generate mockgen -destination=repository/repomock/repo_mock.go -package=repomock -source=repository.go

type AttendanceReaderRepo interface {
	GetByName(ctx context.Context, name string) (entity.Attendance, error)

	WithTx(db.Tx) AttendanceReaderRepo
}

type AttendanceWriterRepo interface {
	Insert(ctx context.Context, attendance entity.Attendance) error

	WithTx(db.Tx) AttendanceWriterRepo
}

type BidReaderRepo interface {
	GetByID(ctx context.Context, id int) (entity.Bid, error)
	GetByIDWithProduct(ctx context.Context, id int) (entity.BidWithProduct, error)
	GetByIDsWithProduct(ctx context.Context, ids []int) ([]entity.BidWithProduct, error)
	GetByUserIDAndProductID(ctx context.Context, userID, productID int) (entity.Bid, error)
	ListUserBiddedProducts(ctx context.Context, userID int, page int, limit int) ([]entity.BidWithProduct, error)
	ListByProductIDAndStatus(ctx context.Context, productID, status, page, limit int) ([]entity.Bid, error)

	WithTx(db.Tx) BidReaderRepo
}

type BidWriterRepo interface {
	Insert(ctx context.Context, bid entity.Bid) (entity.Bid, error)
	UpdateAmount(ctx context.Context, bidID, newAmount int) error
	UpdateStatus(ctx context.Context, bidID, status int) error

	WithTx(db.Tx) BidWriterRepo
}

type BidHistoryReaderRepo interface {
	GetByID(ctx context.Context, id int) (entity.Bid, error)
	ListByProductID(ctx context.Context, productID int, page int, limit int) ([]entity.BidWithBidder, error)

	WithTx(db.Tx) BidHistoryReaderRepo
}

type BidHistoryWriterRepo interface {
	Insert(ctx context.Context, bid entity.Bid) (entity.Bid, error)

	WithTx(db.Tx) BidHistoryWriterRepo
}

type ProductReaderRepo interface {
	GetByIDAndLock(ctx context.Context, productID int) (entity.Product, error)
	GetByIDWithSeller(ctx context.Context, productID int) (entity.ProductWithSeller, error)
	GetProductByOwnerUserID(ctx context.Context, ownerUserId int) ([]entity.Product, error)
	GetFinishedProducts(ctx context.Context, page, limit int) ([]entity.Product, error)

	WithTx(db.Tx) ProductReaderRepo
}

type ProductWriterRepo interface {
	InsertProduct(ctx context.Context, product entity.Product) error
	UpdateLastBidID(ctx context.Context, productID, lastBidID int) error
	UpdateStatus(ctx context.Context, productID, status int) error

	WithTx(db.Tx) ProductWriterRepo
}

type WalletReaderRepo interface {
	GetByUserID(ctx context.Context, userID int) (entity.Wallet, error)

	WithTx(db.Tx) WalletReaderRepo
}

type WalletWriterRepo interface {
	InjectWalletByUserID(ctx context.Context, userID, injection int) error
	DeductWallet(ctx context.Context, id, deduction int) error

	WithTx(db.Tx) WalletWriterRepo
}
