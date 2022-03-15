package ucv1

import (
	"context"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/entity"
	"github.com/palantir/stacktrace"
)

type productUC struct {
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
