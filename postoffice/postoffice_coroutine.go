package postoffice

import (
	"log"

	"github.com/chainpusher/chainpusher/model"
)

type PostOfficeCoroutine struct {
	Transports []Transport
}

func NewPostOfficeCoroutine(transports []Transport) *PostOfficeCoroutine {
	return &PostOfficeCoroutine{
		Transports: transports,
	}
}

func (p *PostOfficeCoroutine) GetTransports() []Transport {
	return p.Transports
}

func (p *PostOfficeCoroutine) Deliver(transactions []*model.Transaction) error {

	log.Println("Delivering transaction", transactions)

	for _, transport := range p.GetTransports() {
		go func(transport Transport) {
			if transport.Deliver(transactions) != nil {
				log.Println("Failed to deliver transaction", transactions)
			}
		}(transport)
	}
	return nil
}
