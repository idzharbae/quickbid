package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter/service/types"
	"github.com/julienschmidt/httprouter"
	"github.com/palantir/stacktrace"
)

type ProductService struct {
	productUC src.ProductUC
}

func NewProductService(productUC src.ProductUC) HttpService {
	return &ProductService{productUC: productUC}
}

func (p *ProductService) GetHandles() []Handle {
	return []Handle{
		{
			Method: http.MethodGet,
			Path:   "/product/:id",
			Handle: p.GetProductHandler,
		},
	}
}

func (p *ProductService) GetProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	productID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid user_id", 400)
		return
	}

	product, err := p.productUC.GetByIDWithSeller(r.Context(), productID)
	if err != nil {
		log.Println(stacktrace.Propagate(err, "[ProductService][GetByIDWithSeller]"))
		http.Error(w, stacktrace.RootCause(err).Error(), 500)
		return
	}

	res := types.GetProductResponseFromProductAndSellerEntity(product)
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Println(stacktrace.Propagate(err, "[ProductService][Marshal]"))
		http.Error(w, stacktrace.RootCause(err).Error(), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(resJson)
}
