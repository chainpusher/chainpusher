package cron

import (
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade"
	"github.com/robfig/cron"
)

type TradePriceJob struct {
	factory facade.PriceServiceFacadeFactory
}

func (job *TradePriceJob) Run() {
	priceServiceFacade := job.factory.NewPriceServiceFacade()
	if err := priceServiceFacade.LoadPrices(); err != nil {

	}
}

func NewTradePriceJob(factory facade.PriceServiceFacadeFactory) *TradePriceJob {
	return &TradePriceJob{factory: factory}
}

func NewCron(jobs ...cron.Job) (*cron.Cron, error) {
	c := cron.New()

	for _, job := range jobs {
		if err := c.AddJob("@hourly", job); err != nil {
			return nil, err
		}
	}

	return c, nil
}
