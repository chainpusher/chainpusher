package config_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
)

func TestIsOnList(t *testing.T) {
	r := config.NewConfigWatchlistRepository([]string{"wallet1", "wallet2"})

	transactions := []*model.Transaction{
		{Payee: "wallet1"},
		{Payer: "wallet2"},
		{Payee: "wallet3"},
	}

	watched := r.In(transactions)
	if len(watched) != 2 {
		t.Error("watched is incorrected", watched)
	}
}
