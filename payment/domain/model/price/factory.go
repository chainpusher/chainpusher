package price

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/counter"
	"strconv"
)

type Factory struct {
	counterRepository counter.Repository

	priceRepository Repository
}

func (f *Factory) AskForPrice(amount int64) (*Price, error) {
	var p *Price
	var c *counter.Counter
	var err error
	key := strconv.FormatInt(amount, 10)
	if p, err = f.priceRepository.FindPriceByAmount(amount); p != nil {
		p.Used = true
		return p, nil
	}

	affected, err := f.counterRepository.IncrementCounterByKey(key)
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		c = counter.NewCounter(key)
		if err = f.counterRepository.Save(c); err != nil {
			return nil, err
		}
	} else {
		if c, err = f.counterRepository.FindCounterByKey(key); err != nil {
			return nil, err
		}
	}

	p = &Price{
		Amount: amount,
		State:  c.Value,
		Used:   true,
	}

	if err := f.priceRepository.Save(p); err != nil {
		return nil, err
	}

	return p, nil
}

func NewFactory(counterRepository counter.Repository, priceRepository Repository) *Factory {
	return &Factory{
		counterRepository: counterRepository,
		priceRepository:   priceRepository,
	}
}
