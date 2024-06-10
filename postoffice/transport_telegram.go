package postoffice

import (
	"net/http"
	"strings"
	"sync"

	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/ethereum/go-ethereum/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var transportTelegram *TransportTelegram
var creating sync.WaitGroup
var called = false

type TelegramBot struct {
	ChatIdentifications []int64

	Client *tgbotapi.BotAPI
}

type TransportTelegram struct {
	Bots map[string]*TelegramBot
}

func NewTelegramBot(token string) (*TelegramBot, error) {

	httpProxy, err := sys.GetEnv("HTTP_PROXY")
	if nil == err {
		log.Debug("HTTP_PROXY: ", httpProxy)
	}
	var bot *tgbotapi.BotAPI

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	bot, err = tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, httpClient)

	if err != nil {
		logrus.Warnf("Failed to create Telegram bot with token %s: %v", token, err)
		return nil, err
	}

	logrus.Debug("Authorized on account ", bot.Self.UserName)
	updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   10,
		Timeout: 10,
	})
	logrus.Debug("Updates: ", updates)

	if err != nil {
		logrus.Warnf("Failed to get updates for Telegram bot with token %s: %v", token, err)
		return nil, err
	}

	var chatIdentifications []int64 = make([]int64, 0)
	for _, update := range updates {
		chatIdentifications = append(chatIdentifications, update.Message.Chat.ID)
	}
	logrus.Debug("Chat identifications: ", chatIdentifications)

	return &TelegramBot{
		ChatIdentifications: chatIdentifications,
		Client:              bot,
	}, nil
}

func NewTransportTelegram(tokens []string) Transport {
	if called {
		logrus.Debug("Did not create Telegram transport because it was already created.")
		creating.Wait()
	}
	called = true
	creating.Add(1)

	if transportTelegram != nil {
		return transportTelegram
	}
	bots := make(map[string]*TelegramBot)

	for _, token := range tokens {
		bot, err := NewTelegramBot(token)
		if err != nil {
			continue
		}

		bots[token] = bot
	}

	logrus.Debug("Created Telegram transport with ", len(bots), " bots")

	transportTelegram = &TransportTelegram{Bots: bots}
	creating.Done()
	return transportTelegram
}

func (t *TransportTelegram) Deliver(transactions []*model.Transaction) error {
	logrus.Debugf("Deliver %d transactions to telegram message", len(transactions))

	messages := make([]string, 0)
	for _, transaction := range transactions {
		messages = append(messages, transaction.SocialMediaMessage())
	}
	message := strings.Join(messages, "\n")

	for _, bot := range t.Bots {
		go func(bot *TelegramBot, message string) {
			defer func() {
				if r := recover(); r != nil {
					logrus.Warnf("Recovered from panic: %v", r)
				}
			}()

			if len(bot.ChatIdentifications) == 0 {
				log.Info("No chat identifications found for Telegram bot")
				return
			}
			telegramMessage := tgbotapi.NewMessage(bot.ChatIdentifications[0], message)
			sent, err := bot.Client.Send(telegramMessage)

			if err != nil {
				logrus.Warnf("Failed to send message %s to Telegram chat %d: %v", message, bot.ChatIdentifications[0], err)
				return
			}

			logrus.Debugf("Message sent to Telegram chat %d: %v", bot.ChatIdentifications[0], sent)
		}(bot, message)
	}
	return nil
}
