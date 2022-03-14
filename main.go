package main

import (
	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/http"
)

func main() {
	quickbidApp := app.NewQuickBidApp("files/config/quickbid.json")
	http.Start(quickbidApp)
}
