package config

import (
	"github.com/chainpusher/blockchain/model"
	model2 "github.com/chainpusher/chainpusher/domain/model"
)

type WatchlistRepository struct {
	Wallets map[string]bool
}

func (cwr *WatchlistRepository) In(block *model.Block) []*model.Transaction {
	watched := make([]*model.Transaction, 0)

	for _, transaction := range block.Transactions {
		var exists bool

		_, exists = cwr.Wallets[transaction.Payer]
		if !exists {
			_, exists = cwr.Wallets[transaction.Payee]
		}
		if !exists {
			continue
		}
		watched = append(watched, transaction)
	}
	return watched
}

func NewConfigWatchlistRepository(wallets []string) model2.WatchlistRepository {
	var walletsMap = make(map[string]bool)

	for _, wallet := range wallets {
		walletsMap[wallet] = true
	}
	return &WatchlistRepository{Wallets: walletsMap}
}
