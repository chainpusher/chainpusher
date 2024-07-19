package commands

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"math/big"
)

type BroadcastBlockListener struct {
	facade facade.TinyBlockServiceFacade
}

func (b BroadcastBlockListener) BeforeQuerying(height *big.Int) {

}

func (b BroadcastBlockListener) AfterRawQuerying(block interface{}, err error) {
}

func (b BroadcastBlockListener) AfterQuerying(block *model.Block, err error) {
	if err != nil {
		return
	}
	b.facade.Broadcast(block)
}

func NewBroadcastBlockListener() service.BlockListener {
	return &BroadcastBlockListener{}
}