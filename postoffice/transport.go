package postoffice

import "github.com/chainpusher/chainpusher/model"

type Transport interface {
	Deliver(transactions []*model.Transaction) error
}
