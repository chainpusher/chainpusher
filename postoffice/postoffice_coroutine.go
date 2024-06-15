package postoffice

import (
	"github.com/chainpusher/blockchain/model"
	"log"

	"github.com/sirupsen/logrus"
)

type Coroutine struct {
	Transports []Transport
}

func NewPostOfficeCoroutine(transports []Transport) *Coroutine {
	return &Coroutine{
		Transports: transports,
	}
}

func (p *Coroutine) GetTransports() []Transport {
	return p.Transports
}

func (p *Coroutine) Deliver(block *model.Block) error {
	for _, transport := range p.GetTransports() {
		go func(transport Transport) {

			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered from panic", r)
				}
			}()

			if transport.Deliver(block) != nil {
				logrus.Errorf("Failed to deliver the block %d", block.Height)
			}
		}(transport)
	}
	return nil
}
