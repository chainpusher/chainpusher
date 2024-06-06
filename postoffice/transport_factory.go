package postoffice

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/sirupsen/logrus"
)

func NewTransportTelegramFromConfig(cfg *config.Config) Transport {
	if cfg.Telegram.Tokens == nil {
		return nil
	}
	logrus.Info("Creating Telegram transport")

	defer logrus.Info("Telegram transport created")
	return NewTransportTelegram(cfg.Telegram.Tokens)
}

func NewTransportHttpFromConfig(cfg *config.Config) Transport {
	if cfg.Http == nil {
		return nil
	}
	logrus.Info("Creating HTTP transport")

	urls := make([]string, 0)
	for _, http := range cfg.Http {
		urls = append(urls, http.Url)
	}

	defer logrus.Info("HTTP transport created")
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
