package httprouter

import (
	"net/http"

	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter/service"

	"github.com/julienschmidt/httprouter"
)

func NewHandler(quickbidApp *app.QuickBid) http.Handler {
	router := httprouter.New()

	services := []service.HttpService{
		service.NewBidService(quickbidApp.UseCases.BidUC),
		service.NewAttendanceService(quickbidApp.UseCases.AttendanceUC),
		service.NewProductService(quickbidApp.UseCases.ProductUC),
	}

	for _, service := range services {
		for _, handle := range service.GetHandles() {
			router.Handle(handle.Method, handle.Path, handle.Handle)
		}
	}

	return router
}
