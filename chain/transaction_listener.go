package chain

import "github.com/chainpusher/chainpusher/model"

type TransactionListener interface {
	OnTransaction(model.Transaction)
}
