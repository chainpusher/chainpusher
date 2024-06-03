package postoffice

import (
	"log"

	"github.com/chainpusher/chainpusher/model"
)

type PostOfficeCoroutine struct {
}

func NewPostOfficeCoroutine() *PostOfficeCoroutine {
	return &PostOfficeCoroutine{}
}

func (p *PostOfficeCoroutine) Deliver(transactions []*model.Transaction) error {

	log.Println("Delivering transaction", transactions)

	for _, transport := range CreateTransportFactory() {
		go func(transport Transport) {
			if transport.Deliver(transactions) != nil {
				log.Println("Failed to deliver transaction", transactions)
			}
		}(transport)
	}
	return nil
}
