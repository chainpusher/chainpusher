package application

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/config"
	model2 "github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
)

type DefaultAnalysisService struct {
	WatchlistRepository model2.WatchlistRepository

	postOffice postoffice.PostOffice
}

func NewDefaultAnalysisService(cfg *config.Config) *DefaultAnalysisService {
	watchlist := config.NewConfigWatchlistRepository(cfg.Wallets)
	t := postoffice.NewTransportFromConfig(cfg)
	ps := postoffice.NewPostOfficeCoroutine(t)

	return &DefaultAnalysisService{
		WatchlistRepository: watchlist,
		postOffice:          ps,
	}
}

func NewSimpleDefaultAnalysisService(wr model2.WatchlistRepository, po postoffice.PostOffice) *DefaultAnalysisService {
	return &DefaultAnalysisService{
		WatchlistRepository: wr,
		postOffice:          po,
	}
}

func (t *DefaultAnalysisService) AnalyzeTrade(block *model.Block) error {
	if watched := t.WatchlistRepository.In(block); len(watched) > 0 {
		aBlock := block.CloneWithTransactions(block.Transactions)
		err := t.postOffice.Deliver(aBlock)

		if err != nil {
			// TODO: handle error
		}
	}

	return nil
}
