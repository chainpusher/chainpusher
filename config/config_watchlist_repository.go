package config

import (
	"github.com/chainpusher/chainpusher/model"
)

type ConfigWatchlistRepository struct {
	Wallets map[string]bool
}

func (cwr *ConfigWatchlistRepository) IsOnList(address string) bool {
	watched, exists := cwr.Wallets[address]
	if !exists {
		return false
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
