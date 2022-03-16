package main

import (
	"flag"

	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/cron"
	"github.com/idzharbae/quickbid/src/delivery/http"
)

func main() {
	service := flag.String("service", "http", "Type of service (http/cron)")
	flag.Parse()

	quickbidApp := app.NewQuickBidApp("files/config/quickbid.json")
	switch *service {
	case "cron":
		cron.Start(quickbidApp)
	default:
		http.Start(quickbidApp)
	}

}
