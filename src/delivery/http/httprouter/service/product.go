package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter/service/types"
	"github.com/idzharbae/quickbid/src/requests"
	"github.com/julienschmidt/httprouter"
	"github.com/palantir/stacktrace"
)

type ProductService struct {
	productUC src.ProductUC
}

func NewProductService(productUC src.ProductUC) HttpService {
	return &ProductService{productUC: productUC}
}

type (
	UploadProductRequest struct {
		Name         string    `json:"name"`
		InitialPrice int       `json:"initial_price"`
		StartBidDate time.Time `json:"start_time"`
		EndBidDate   time.Time `json:"end_time"`
		OwnerUserID  int       `json:"seller_id"`
		ImageURL     string    `json:"image_url"`
		BidIncrement int       `json:"bid_increment"`
	}

	GetProductByOwnerUserIDRequest struct {
		OwnerUserID int `json:"owner_user_id"`
	}
)

func (ps *ProductService) GetHandles() []Handle {
	return []Handle{
		{
			Method: http.MethodPost,
			Path:   "/product/add",
			Handle: ps.UploadProductHandler,
		},
		{
			Method: http.MethodPost,
			Path:   "/product/get/by_seller_id",
			Handle: ps.GetProductByOwnerUserIDHandler,
		},
		{
			Method: http.MethodGet,
			Path:   "/product/:id",
			Handle: ps.GetProductHandler,
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
	}
	w.Write(resJson)
}

func (ps *ProductService) UploadProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	var req UploadProductRequest
	err = json.Unmarshal([]byte(data), &req)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	err = ps.productUC.UploadProduct(r.Context(), parseUploadProductRequest(req))
	if err != nil {
		log.Printf("[UploadProductHandler][productUC] Product: %s\n", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write([]byte("Success add product!"))
}

func (ps *ProductService) GetProductByOwnerUserIDHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	var req GetProductByOwnerUserIDRequest
	err = json.Unmarshal([]byte(data), &req)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	resp, err := ps.productUC.GetProductByOwnerUserID(r.Context(), parseGetProductByOwnerUserIDRequest(req))
	if err != nil {
		log.Printf("[GetProductByOwnerUserIDHandler][productUC] Error: %s\n", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[GetProductByOwnerUserIDHandler][productUC] Failed Marshall Error: %s\n", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(respJson)
}

func parseUploadProductRequest(req UploadProductRequest) requests.UploadProductRequest {
	return requests.UploadProductRequest{
		Name:         req.Name,
		InitialPrice: req.InitialPrice,
		StartBidDate: req.StartBidDate,
		EndBidDate:   req.EndBidDate,
		OwnerUserID:  req.OwnerUserID,
		ImageURL:     req.ImageURL,
		BidIncrement: req.BidIncrement,
	}
}

func parseGetProductByOwnerUserIDRequest(req GetProductByOwnerUserIDRequest) requests.GetProductByOwnerUserIDRequest {
	return requests.GetProductByOwnerUserIDRequest{
		OwnerUserID: req.OwnerUserID,
	}
}
