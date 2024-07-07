package model

import "github.com/chainpusher/blockchain/model"

type TradingListener interface {
	BlockGenerated(block *model.Block)
}
