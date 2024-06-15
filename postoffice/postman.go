package postoffice

import (
	"github.com/chainpusher/blockchain/model"
)

type Postman interface {
	Deliver(block *model.Block) error
}
