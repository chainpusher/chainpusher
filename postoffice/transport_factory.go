package postoffice

import (
	"github.com/chainpusher/chainpusher/config"
)

func NewTransportTelegramFromConfig(cfg *config.Config) Transport {
	if cfg.Telegram.Tokens == nil {
		return nil
	}
	return NewTransportTelegram(cfg.Telegram.Tokens)
}

func NewTransportHttpFromConfig(cfg *config.Config) Transport {
	if cfg.Http == nil {
		return nil
	}

	urls := make([]string, 0)
	for _, http := range cfg.Http {
		urls = append(urls, http.Url)
	}

	return NewTransportHttp(urls)
}

func NewTransportFromConfig(cfg *config.Config) []Transport {
	transports := make([]Transport, 0)
	telegram := NewTransportTelegramFromConfig(cfg)
	if telegram != nil {
		transports = append(transports, telegram)
	}

	http := NewTransportHttpFromConfig(cfg)
	if http != nil {
		transports = append(transports, http)
	}

	return transports
}
