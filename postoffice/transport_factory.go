package postoffice

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/sirupsen/logrus"
)

func CreateTelegramTransport(cfg *config.Config) Transport {
	if cfg.Telegram.Tokens == nil {
		return nil
	}
	logrus.Info("Creating Telegram transport")

	defer logrus.Info("Telegram transport created")
	return NewTransportTelegram(cfg.Telegram.Tokens)
}

func CreateTransportFactory(cfg *config.Config) []Transport {
	transports := make([]Transport, 0)
	telegram := CreateTelegramTransport(cfg)
	if telegram != nil {
		transports = append(transports, telegram)
	}
	return transports
}
