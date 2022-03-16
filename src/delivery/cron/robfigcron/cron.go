package robfigcron

import (
	"github.com/idzharbae/quickbid/src/app"
	"github.com/idzharbae/quickbid/src/delivery/cron/jobs"
	"github.com/robfig/cron/v3"
)

func Run(qbApp *app.QuickBid) {
	c := cron.New(cron.WithSeconds())
	c.AddJob("@every 1m", jobs.NewBidFinisher(qbApp.UseCases.BidUC, qbApp.UseCases.ProductUC))

	c.Run()

	// job := jobs.NewBidFinisher(qbApp.UseCases.BidUC, qbApp.UseCases.ProductUC)
	// job.Run()
}
