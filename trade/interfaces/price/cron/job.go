package cron

import (
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade"
	"github.com/robfig/cron"
)

type TradePriceJob struct {
	priceServiceFacade facade.ServiceFacade
}

func (job *TradePriceJob) Run() {
	if err := job.priceServiceFacade.LoadPrices(); err != nil {

	}
}

func NewTradePriceJob(priceServiceFacade facade.ServiceFacade) *TradePriceJob {
	return &TradePriceJob{priceServiceFacade: priceServiceFacade}
}

func NewCron(jobs []cron.Job) (*cron.Cron, error) {
	c := cron.New()
	for _, job := range jobs {
		if err := c.AddJob("@hourly", job); err != nil {
			return nil, err
		}
	}

	return c, nil
}
