package application_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
)

func TestTransactionServiceAnalyzeTrade(t *testing.T) {
	transactions := []*model.Transaction{
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ethereum,
			Payee:          "payee",
			Payer:          "payer",
			Amount:         *big.NewInt(100),
		},
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ethereum,
			Payee:          "payee2",
			Payer:          "payer2",
			Amount:         *big.NewInt(200),
		},
	}

	repository := &WatchlistRepositoryMock{called: false}
	transportMock := &TransportMock{called: false}
	transports := []postoffice.Transport{transportMock}
	ps := &postoffice.PostOfficeCoroutine{Transports: transports}
	service := &application.TransactionService{
		WatchlistRepository: repository,
		Postoffice:          ps,
	}
	service.AnalyzeTrade(transactions)
	time.Sleep(1 * time.Second)

	if !repository.called || !transportMock.called {
		t.Error("Failed to call repository or transport")
		return
	}
	t.Log("Transaction service analyzed trade successfully")
}

type WatchlistRepositoryMock struct {
	called bool
}

func (w *WatchlistRepositoryMock) In(transactions []*model.Transaction) []*model.Transaction {
	w.called = true
	return transactions
}

type TransportMock struct {
	called bool
}

func (t *TransportMock) Deliver(transactions []*model.Transaction) error {
	t.called = true
	return nil
}
