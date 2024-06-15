package postoffice

import "github.com/chainpusher/blockchain/model"

type PostOffice interface {
	Deliver(block *model.Block) error

	GetTransports() []Transport
}
