package ucv1

import (
	"context"
	"time"

	"github.com/ggwhite/go-masker"
	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/bridge/db"
	"github.com/idzharbae/quickbid/src/bridge/transactioner"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"
	"github.com/jackc/pgx/v4"

	"github.com/palantir/stacktrace"
)

type bidUC struct {
	bidReader        src.BidReaderRepo
	bidWriter        src.BidWriterRepo
	bidHistoryReader src.BidHistoryReaderRepo
	bidHistoryWriter src.BidHistoryWriterRepo
	productReader    src.ProductReaderRepo
	productWriter    src.ProductWriterRepo
	walletReader     src.WalletReaderRepo
	walletWriter     src.WalletWriterRepo

	txner transactioner.Transactioner
}

func NewBidUC(bidReader src.BidReaderRepo,
	bidWriter src.BidWriterRepo,
	bidHistoryReader src.BidHistoryReaderRepo,
	bidHistoryWriter src.BidHistoryWriterRepo,
	productReader src.ProductReaderRepo,
	productWriter src.ProductWriterRepo,
	walletReader src.WalletReaderRepo,
	walletWriter src.WalletWriterRepo,
	txner transactioner.Transactioner) src.BidUC {
	return &bidUC{
		bidReader:        bidReader,
		bidWriter:        bidWriter,
		bidHistoryReader: bidHistoryReader,
		bidHistoryWriter: bidHistoryWriter,
		productReader:    productReader,
		productWriter:    productWriter,
		walletReader:     walletReader,
		walletWriter:     walletWriter,
		txner:            txner,
	}
}

func (b *bidUC) ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.BidWithProduct, error) {
	bids, err := b.bidReader.ListUserBiddedProducts(ctx, req.UserID, req.Page, req.Limit)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidUC][ListUserBiddedProducts]")
	}

	return bids, nil
}

func (b *bidUC) ListByProduct(ctx context.Context, req requests.ListBidsByProductRequest) ([]entity.BidWithBidder, error) {
	bids, err := b.bidHistoryReader.ListByProductID(ctx, req.ProductID, req.Page, req.Limit)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[bidUC][ListByProduct]")
	}

	// Mask bidder name
	for i := range bids {
		bids[i].Bidder.Name = masker.Name(bids[i].Bidder.Name)
	}

	return bids, nil
}

func (b *bidUC) BidProduct(ctx context.Context, req requests.BidProductRequest) (entity.Bid, error) {
	var newBid entity.Bid

	txnErr := b.txner.DoWithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		bidReader := b.bidReader.WithTx(tx)
		bidWriter := b.bidWriter.WithTx(tx)
		bidHistoryReader := b.bidHistoryReader.WithTx(tx)
		bidHistoryWriter := b.bidHistoryWriter.WithTx(tx)
		productReader := b.productReader.WithTx(tx)
		productWriter := b.productWriter.WithTx(tx)
		walletReader := b.walletReader.WithTx(tx)
		walletWriter := b.walletWriter.WithTx(tx)

		validateRequest := func(req requests.BidProductRequest, product entity.Product) error {
			now := time.Now()
			if now.Before(product.StartBidDate) || now.After(product.EndBidDate) {
				return stacktrace.NewError("can't bid outside product's bidding time range: %v - %v", product.StartBidDate, product.EndBidDate)
			}

			if req.Amount < product.InitialPrice {
				return stacktrace.NewError("bid amount is lower than minimum amount")
			}

			if req.Amount%product.BidIncrement != 0 {
				return stacktrace.NewError("bid amount is not divisible by bid increment")
			}

			return nil
		}

		deductWallet := func(req requests.BidProductRequest, bidderLastBid entity.Bid) error {
			walletDeduction := req.Amount - bidderLastBid.Amount
			wallet, err := walletReader.GetByUserID(ctx, req.UserID)
			if err != nil {
				return stacktrace.Propagate(err, "[bidUC][BidProduct][walletReader][GetByUserID]")
			}

			if wallet.Amount < walletDeduction {
				return stacktrace.NewError("not enough balance")
			}

			err = walletWriter.DeductWallet(ctx, wallet.ID, walletDeduction)
			if err != nil {
				stacktrace.Propagate(err, "[bidUC][BidProduct][walletWriter][DeductWallet]")
			}

			return nil
		}

		upsertNewBid := func(req requests.BidProductRequest, bidderLastBidID int) (entity.Bid, error) {
			var err error

			newBid := entity.Bid{
				ID:        bidderLastBidID,
				UserID:    req.UserID,
				ProductID: req.ProductID,
				Amount:    req.Amount,
				BidTime:   time.Now(),
			}
			if bidderLastBidID > 0 {
				err := bidWriter.UpdateAmount(ctx, bidderLastBidID, req.Amount)
				if err != nil {
					stacktrace.Propagate(err, "[bidUC][BidProduct][bidWriter][UpdateAmount]")
				}

				return newBid, nil
			}

			newBid, err = bidWriter.Insert(ctx, newBid)
			if err != nil {
				stacktrace.Propagate(err, "[bidUC][BidProduct][bidWriter][UpdateAmount]")
			}

			return newBid, nil
		}

		product, err := productReader.GetByIDAndLock(ctx, req.ProductID)
		if err != nil {
			return stacktrace.Propagate(err, "[bidUC][BidProduct][productReader][GetByIDAndLock]")
		}

		err = validateRequest(req, product)
		if err != nil {
			return err
		}

		if product.LastBidID > 0 {
			lastBid, err := bidHistoryReader.GetByID(ctx, product.LastBidID)
			if err != nil {
				return stacktrace.Propagate(err, "[bidUC][BidProduct][bidHistoryReader][GetByID]")
			}

			if req.Amount <= lastBid.Amount {
				return stacktrace.NewError("bid amount is not higher than last bid")
			}

			err = bidWriter.UpdateStatus(ctx, lastBid.ID, entity.BidStatusInactive)
			if err != nil {
				return stacktrace.Propagate(err, "[bidUC][BidProduct][bidWriter][UpdateStatus]")
			}
		}

		bidderLastBid, err := bidReader.GetByUserIDAndProductID(ctx, req.UserID, req.ProductID)
		if err != nil && stacktrace.RootCause(err) != pgx.ErrNoRows {
			return stacktrace.Propagate(err, "[bidUC][BidProduct][bidReader][GetByUserIDAndProductID]")
		}

		err = deductWallet(req, bidderLastBid)
		if err != nil {
			return err
		}

		newBid, err = upsertNewBid(req, bidderLastBid.ID)
		if err != nil {
			return err
		}

		newBidHistory, err := bidHistoryWriter.Insert(ctx, newBid)
		if err != nil {
			stacktrace.Propagate(err, "[bidUC][BidProduct][bidHistoryWriter][Insert]")
		}

		err = productWriter.UpdateLastBidID(ctx, req.ProductID, newBidHistory.ID)
		if err != nil {
			stacktrace.Propagate(err, "[bidUC][BidProduct][productWriter][UpdateLastBidID]")
		}

		return nil
	})
	if txnErr != nil {
		return entity.Bid{}, stacktrace.Propagate(txnErr, "[bidUC][txner][DoWithTx]")
	}

	return newBid, nil
}
