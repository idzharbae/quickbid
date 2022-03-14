package httprouter

import (
	"net/http"

	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter/service"

	"github.com/julienschmidt/httprouter"
)

func NewHandler(quickbidApp *app.QuickBid) http.Handler {
	router := httprouter.New()

	attendanceService := service.NewAttendanceService(quickbidApp.UseCases.AttendanceUC)
	for _, handle := range attendanceService.GetHandles() {
		router.Handle(handle.Method, handle.Path, handle.Handle)
	}

	return router
}
