package application

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
)

type DefaultTransactionService struct {
	WatchlistRepository model.WatchlistRepository

	Postoffice postoffice.PostOffice
}

func NewDefaultTransactionService(cfg *config.Config) *DefaultTransactionService {
	watchlist := config.NewConfigWatchlistRepository(cfg.Wallets)
	t := postoffice.NewTransportFromConfig(cfg)
	ps := postoffice.NewPostOfficeCoroutine(t)

	return &DefaultTransactionService{
		WatchlistRepository: watchlist,
		Postoffice:          ps,
	}
}

func (t *DefaultTransactionService) AnalyzeTrade(transactions []*model.Transaction) error {
	if watched := t.WatchlistRepository.In(transactions); len(watched) > 0 {
		t.Postoffice.Deliver(watched)
	}

	return nil
}
