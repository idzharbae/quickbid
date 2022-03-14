package http

import (
	"log"
	"net/http"

	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/http/httprouter"
)

func Start(qbApp *app.QuickBid) {
	router := httprouter.NewHandler(qbApp)

	log.Println("HTTP LISTENING TO PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
