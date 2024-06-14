package monitor

import (
	"math/big"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/model"
)

type PlatformWatcherV2 struct {
	watcher *Watcher

	blockChainService chain.BlockchainService

	transactionService application.TransactionService

	counter *big.Int
}

func (p *PlatformWatcherV2) Start() {
	p.watcher.StartContinuousMovement()
}

func (w *PlatformWatcherV2) Tick(time time.Time) {
	block := w.FetchingBlock()

	w.InspectionBlockSchedule(block)

	if block.IsAcrossMultipleBlocks() {
		// Start new worker to handle these blocks
	}

}

func (w *PlatformWatcherV2) FetchingBlock() *model.Block {
	var block *model.Block
	var err error

	if w.counter.Int64() == 0 {
		block, err = w.blockChainService.GetLatestBlock()
		w.counter.Set(block.Height)
	} else {
		block, err = w.blockChainService.GetBlock(w.counter)
	}
	w.counter.Add(w.counter, big.NewInt(1))
	if err != nil {
		return nil
	}

	return block
}

func (w *PlatformWatcherV2) InspectionBlockSchedule(block *model.Block) {
	var schedule time.Time = w.watcher.GetCurrentScheduledTime()
	var current time.Time = block.GenerateTimeToNextBlock() // predict next block time of current block
	w.counter.Add(w.counter, big.NewInt(1))
	var predictionError time.Duration = current.Sub(schedule)

	if predictionError < 2*time.Second {
		return
	}

	w.BlockAheadOfSchedule(current)

}

func (w *PlatformWatcherV2) BlockAheadOfSchedule(current time.Time) {
	w.watcher.AdvanceTimeTo(current)
}

func (w *PlatformWatcherV2) BlockWasDelayed() {

}

func (w *PlatformWatcherV2) Stop() {
	// w.watcher.Stop()
}

func NewPlatformWatcherV2(
	watcher *Watcher,
	blockChainService chain.BlockchainService,
	transactionService application.TransactionService,
) *PlatformWatcherV2 {
	return &PlatformWatcherV2{
		watcher:            watcher,
		blockChainService:  blockChainService,
		transactionService: transactionService,
		counter:            big.NewInt(0),
	}
}
