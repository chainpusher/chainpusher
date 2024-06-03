package postoffice

import (
	"github.com/chainpusher/chainpusher/model"
)

type Postman interface {
	Deliver(transactions []*model.Transaction) error
}
