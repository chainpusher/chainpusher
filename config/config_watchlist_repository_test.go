package config_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/config"
)

func TestIsOnList(t *testing.T) {
	r := config.NewConfigWatchlistRepository([]string{"wallet1", "wallet2"})
	if !r.IsOnList("wallet1") {
		t.Error("Expected wallet1 to be on the list")
	}

	if r.IsOnList("wallet3") {
		t.Error("Expected wallet3 to not be on the list")
	}

	if !r.IsOnList("wallet2") {
		t.Error("Expected wallet2 to not be on the list")
	}
}
