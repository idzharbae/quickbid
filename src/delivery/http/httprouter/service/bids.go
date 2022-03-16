package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/idzharbae/quickbid/src"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter/service/types"
	"github.com/idzharbae/quickbid/src/requests"

	"github.com/julienschmidt/httprouter"
	"github.com/palantir/stacktrace"
)

type BidService struct {
	BidUC src.BidUC
}

func NewBidService(BidUC src.BidUC) *BidService {
	return &BidService{BidUC: BidUC}
}

func (bs *BidService) GetHandles() []Handle {
	return []Handle{
		{
			Method: http.MethodGet,
			Path:   "/api/bids/user/:user_id",
			Handle: bs.ListUserBidsHandler,
		},
		{
			Method: http.MethodGet,
			Path:   "/api/bids/product/:product_id",
			Handle: bs.ListProductBidsHandler,
		},
	}
}

func (bs *BidService) ListUserBidsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Invalid user_id", 400)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "Invalid limit", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page", 400)
		return
	}

	bids, err := bs.BidUC.ListUserBiddedProducts(r.Context(), requests.ListUserBiddedProductsRequest{
		UserID: userId,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		log.Println(stacktrace.Propagate(err, "[BidService][ListUserBidsHandler]"))
		http.Error(w, stacktrace.RootCause(err).Error(), 500)
		return
	}

	bidsResponse := types.ListBidResponseFromBidsWithProduct(bids)
	bidsResponseJson, err := json.Marshal(bidsResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(bidsResponseJson)
}

func (bs *BidService) ListProductBidsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	productID, err := strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		http.Error(w, "Invalid product_id", 400)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "Invalid limit", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page", 400)
		return
	}

	bids, err := bs.BidUC.ListByProduct(r.Context(), requests.ListBidsByProductRequest{
		ProductID: productID,
		Page:      page,
		Limit:     limit,
	})
	if err != nil {
		log.Println(stacktrace.Propagate(err, "[BidService][ListByProduct]"))
		http.Error(w, stacktrace.RootCause(err).Error(), 500)
		return
	}

	bidsResponse := types.ListProductBidsResponseFromEntity(bids)
	bidsResponseJson, err := json.Marshal(bidsResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(bidsResponseJson)
}
