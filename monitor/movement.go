package monitor

import (
	"github.com/chainpusher/blockchain/model"
	"time"
)

type Movement interface {
	BeforeQueryingBlock()

	AfterQueryingBlock(block *model.Block)

	WaitTheNextBlockToBeGenerated(watcher *PlatformWatcher, block *model.Block)
}

type DefaultMovement struct {
}

func (d *DefaultMovement) BeforeQueryingBlock() {

}

func (d *DefaultMovement) AfterQueryingBlock(_ *model.Block) {

}

func (d *DefaultMovement) WaitTheNextBlockToBeGenerated(watcher *PlatformWatcher, _ *model.Block) {
	time.Sleep(watcher.GetTimeForBlockGenerated() * time.Second)
}
