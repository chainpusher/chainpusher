package postoffice

import "github.com/chainpusher/chainpusher/model"

type Transport interface {
	Deliver(transaction *model.Transaction) error
}
