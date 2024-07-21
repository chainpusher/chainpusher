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

func (b BroadcastBlockListener) BeforeQuerying(_ *big.Int) {

}

func (b BroadcastBlockListener) AfterRawQuerying(_ interface{}, _ error) {
}

func (b BroadcastBlockListener) AfterQuerying(block *model.Block, err error) {
	if err != nil {
		return
	}
	b.facade.Broadcast(block)
}

func NewBroadcastBlockListener(theFacade facade.TinyBlockServiceFacade) service.BlockListener {
	return &BroadcastBlockListener{facade: theFacade}
}
