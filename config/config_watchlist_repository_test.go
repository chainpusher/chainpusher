package config_test

import (
	"github.com/chainpusher/blockchain/model"
	"testing"

	"github.com/chainpusher/chainpusher/config"
)

func TestIsOnList(t *testing.T) {
	r := config.NewConfigWatchlistRepository([]string{"wallet1", "wallet2"})

	transactions := []*model.Transaction{
		{Payee: "wallet1"},
		{Payer: "wallet2"},
		{Payee: "wallet3"},
	}

	watched := r.In(&model.Block{Transactions: transactions})
	if len(watched) != 2 {
		t.Error("watched is incorrected", watched)
	}
}
