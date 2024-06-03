package config

import (
	"github.com/chainpusher/chainpusher/model"
)

type ConfigWatchlistRepository struct {
	Wallets map[string]bool
}

func (cwr *ConfigWatchlistRepository) In(transactions []*model.Transaction) []*model.Transaction {
	watched := make([]*model.Transaction, 0)

	for _, transaction := range transactions {
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

func NewConfigWatchlistRepository(wallets []string) model.WatchlistRepository {
	var walletsMap map[string]bool = make(map[string]bool)

	for _, wallet := range wallets {
		walletsMap[wallet] = true
	}
	return &ConfigWatchlistRepository{Wallets: walletsMap}
}
