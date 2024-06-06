package config_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/config"
)

func TestParseConfigFromYamlText(t *testing.T) {
	var text string = `wallets:
  - a
  - b

telegram:
  token: 7041850206:AAEUECtq8CwY-jTJ5SJIdE3Q2hzuWLWcW3s

http:
  - url: https://httpbin.org/post

logger:
  level: debug`
	cfg, err := config.ParseConfigFromYamlText(text)
	if err != nil {
		t.Error(err)
	}

	if len(cfg.Wallets) != 2 {
		t.Error("Expected 2 wallets")
	}

	if cfg.Http[0].Url != "https://httpbin.org/post" {
		t.Error("Expected https://httpbin.org/post")
	}
}
