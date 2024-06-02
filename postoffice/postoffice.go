package postoffice

import "github.com/chainpusher/chainpusher/model"

type PostOffice interface {
	Deliver(transaction *model.Transaction) error
}
