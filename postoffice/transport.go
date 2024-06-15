package postoffice

import "github.com/chainpusher/blockchain/model"

type Transport interface {
	Deliver(block *model.Block) error
}
