package facade

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
)

type TinyBlockServiceFacade interface {
	GetTransactions(command *dto.QueryTransactionsCommand) ([]*model.Transaction, error)

	Subscribe(clientId int64)

	Unsubscribe(clientId int64)

	Broadcast(block *model.Block)
}
