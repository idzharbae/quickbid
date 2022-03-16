package ucv1

import (
	"context"
	"log"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/idzharbae/quickbid/src/requests"
	"github.com/palantir/stacktrace"
)

type productUC struct {
	productWriter src.ProductWriterRepo
	productReader src.ProductReaderRepo
}

func NewProductUC(productReader src.ProductReaderRepo) src.ProductUC {
	return &productUC{productReader: productReader}
}

func (p *productUC) GetByIDWithSeller(ctx context.Context, productID int) (entity.ProductWithSeller, error) {
	product, err := p.productReader.GetByIDWithSeller(ctx, productID)
	if err != nil {
		return entity.ProductWithSeller{}, stacktrace.Propagate(err, "[productUC][GetByIDWithSeller]")
	}

	return product, nil
}

func (p *productUC) UploadProduct(ctx context.Context, req requests.UploadProductRequest) error {
	product := p.constructProductData(req)
	err := p.productWriter.InsertProduct(ctx, product)
	if err != nil {
		log.Println("failed UploadProduct")
		return err
	}
	return nil
}

func (p *productUC) GetProductByOwnerUserID(ctx context.Context, req requests.GetProductByOwnerUserIDRequest) ([]entity.Product, error) {
	products, err := p.productReader.GetProductByOwnerUserID(ctx, req.OwnerUserID)
	if err != nil {
		log.Println("failed GetProductByOwnerUserID")
		return []entity.Product{}, err
	}
	return products, nil
}

func (p *productUC) constructProductData(req requests.UploadProductRequest) entity.Product {
	return entity.Product{
		Name:         req.Name,
		InitialPrice: req.InitialPrice,
		StartBidDate: req.StartBidDate,
		EndBidDate:   req.EndBidDate,
		OwnerUserID:  req.OwnerUserID,
		ImageURL:     req.ImageURL,
		BidIncrement: req.BidIncrement,
		Status:       entity.ProductStatusActive,
	}
}
