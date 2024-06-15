package application_test

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/postoffice"
)

func TestAnalysisService_AnalyzeTrade(t *testing.T) {
	transactions := []*model.Transaction{
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ether,
			Payee:          "payee",
			Payer:          "payer",
			Amount:         *big.NewInt(100),
		},
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ether,
			Payee:          "payee2",
			Payer:          "payer2",
			Amount:         *big.NewInt(200),
		},
	}

	block := &model.Block{Height: big.NewInt(1), Transactions: transactions}

	repository := &WatchlistRepositoryMock{called: false}
	transportMock := &TransportMock{called: false}
	var transports []postoffice.Transport
	ps := &postoffice.Coroutine{Transports: transports}
	service := application.NewSimpleDefaultAnalysisService(repository, ps)
	err := service.AnalyzeTrade(block)
	assert.Error(t, err, "Failed to analyze trade")
	time.Sleep(1 * time.Second)

	assert.Falsef(t, !repository.called || !transportMock.called, "Failed to call repository or transport")
}

type WatchlistRepositoryMock struct {
	called bool
}

func (w *WatchlistRepositoryMock) In(block *model.Block) []*model.Transaction {
	w.called = true
	return block.Transactions
}

type TransportMock struct {
	called bool
}

func (t *TransportMock) Deliver(_ []*model.Transaction) error {
	t.called = true
	return nil
}
