package price

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/counter"
	"strconv"
	"time"
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
	if p, err = f.priceRepository.FindPriceByAmount(amount); err == nil {
		p.Using = true
		return p, nil
	}

	if c, err = f.counterRepository.FindCounterByKey(key); err != nil {
		return nil, err
	}

	if c == nil {
		c = &counter.Counter{
			Value:     0,
			Key:       key,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	} else {
		c.Value++
	}

	p.State = c.Value
	p.Amount = amount

	if err = f.counterRepository.Save(c); err != nil {
		return nil, err
	}

	return p, nil
}
