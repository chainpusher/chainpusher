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

func (p *PostOfficeCoroutine) Deliver(transaction *model.Transaction) error {

	log.Println("Delivering transaction", transaction)

	for _, transport := range CreateTransportFactory() {
		go func(transport Transport) {
			if transport.Deliver(transaction) != nil {
				log.Println("Failed to deliver transaction", transaction)
			}
		}(transport)
	}
	return nil
}
