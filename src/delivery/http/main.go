package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter"
)

func Start(qbApp *app.QuickBid) {
	router := httprouter.NewHandler(qbApp)

	log.Printf("HTTP LISTENING TO PORT %d\n", qbApp.Cfg.App.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", qbApp.Cfg.App.Port), router))
}
