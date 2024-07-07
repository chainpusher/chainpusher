package model

import "github.com/chainpusher/blockchain/model"

type WatchlistRepository interface {
	In(block *model.Block) []*model.Transaction
}
