package postoffice

import (
	"github.com/chainpusher/chainpusher/model"
)

type Postman interface {
	Deliver(transaction *model.Transaction) error
}