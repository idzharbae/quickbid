package cron

import (
	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/cron/robfigcron"
)

func Start(qbApp *app.QuickBid) {
	robfigcron.Run(qbApp)
}
