package monitor

import (
	"github.com/chainpusher/blockchain/model"
	"math/big"
	"runtime"
	"time"

	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/config"
	"github.com/sirupsen/logrus"
)

type PlatformWatcher struct {
	config                *config.Config
	done                  chan bool
	service               service.BlockChainService
	applicationService    application.AnalysisService
	number                *big.Int
	once                  bool
	isRestart             bool
	timeForBlockGenerated time.Duration
	platform              model.Platform
	movement              Movement
}

func (p *PlatformWatcher) Start() {
	logrus.Infof("Watching %s platform ...", p.platform.String())

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("%s: Recovered: %v", p.platform.String(), r)
			buf := make([]byte, 1024)
			stackSize := runtime.Stack(buf, false)
			logrus.Tracef("%s: Stack trace size: %d", p.platform.String(), stackSize)
			logrus.Errorf("%s: Stack trace: %s", p.platform.String(), string(buf[:stackSize]))
		}
	}()

	p.WatchBlocks()
}

func (p *PlatformWatcher) FetchBlocks() (*model.Block, error) {
	logrus.Debugf("%s: Fetching block %d.", p.platform.String(), p.number)

	var block *model.Block
	var err error
	var height *big.Int = p.number

	if p.number.Cmp(big.NewInt(1)) == -1 {
		block, err = p.service.GetLatestBlock()
	} else {
		block, err = p.RunUntilNothingIsNotFound(height)
	}

	if err != nil {
		return block, err
	}

	logrus.Debugf("%s: Block %d fetched (at %v) with %d transactions", p.platform.String(),
		p.number, block.CreatedAt, len(block.Transactions))

	p.number = block.Height.Add(block.Height, big.NewInt(1))
	err = p.applicationService.AnalyzeTrade(block)

	return block, err
}

func (p *PlatformWatcher) RunUntilNothingIsNotFound(height *big.Int) (*model.Block, error) {
	for i := 0; i < 3; i++ {
		block, err := p.service.GetBlock(height)
		if err != nil {
			logrus.Tracef("%s: Error fetching block %d: %v", p.platform.String(), height, err)
			if service.IsNotFound(err) {
				logrus.Tracef("%s: Block %d not found. Retrying %d times ...", p.platform.String(), height, i+1)
				time.Sleep(1 * time.Second)
				continue
			}

			return nil, err
		}
		return block, nil
	}
	logrus.Warnf("%s: Max retries reached", p.platform.String())
	return nil, NewWatcherError(MaxRetries)
}

func (p *PlatformWatcher) WatchBlocks() {
	for {
		select {
		case <-p.done:
			return
		default:
			p.movement.BeforeQueryingBlock()
			block, err := p.FetchBlocks()

			p.movement.AfterQueryingBlock(block)
			if err != nil {
				logrus.Errorf("%s: Error fetching block: %v", p.platform.String(), err)

				p.Restart()
				continue
			}

			if p.once {
				return
			}

			p.movement.WaitTheNextBlockToBeGenerated(p, block)
		}

	}
}

func (p *PlatformWatcher) Stop() {
	p.done <- true
}

func (p *PlatformWatcher) Restart() {
	logrus.Infof("%s: Restarting watcher ...", p.platform.String())
	p.number = big.NewInt(-1)
}

func (p *PlatformWatcher) GetTimeForBlockGenerated() time.Duration {
	return p.timeForBlockGenerated
}
