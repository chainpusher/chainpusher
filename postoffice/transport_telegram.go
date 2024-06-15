package postoffice

import (
	"github.com/chainpusher/blockchain/model"
	"net/http"
	"strings"
	"sync"

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

	Token string

	WasUpdated bool
}

type TransportTelegram struct {
	Bots map[string]*TelegramBot
}

func NewTelegramBot(token interface{}) (*TelegramBot, error) {
	var textToken string
	var tokenIsObject bool = false
	wasUpdated := false

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

	switch t := token.(type) {
	case string:
		textToken = t
	case map[string]interface{}:
		textToken = t["token"].(string)
		tokenIsObject = true
	}

	bot, err = tgbotapi.NewBotAPIWithClient(textToken, tgbotapi.APIEndpoint, httpClient)

	if err != nil {
		logrus.Warnf("Failed to create Telegram bot with token (%s): %v", textToken, err)
		return nil, err
	}

	logrus.Debug("Authorized on account ", bot.Self.UserName)

	var chatIdentifications []int64 = make([]int64, 0)

	if tokenIsObject && token.(map[string]interface{})["chat_id"] != nil {

		switch anyChat := token.(map[string]interface{})["chat_id"]; anyChat.(type) {
		case int:
			chatIdentifications = append(chatIdentifications, int64(anyChat.(int)))
		case []interface{}:
			for _, chat := range anyChat.([]interface{}) {
				chatIdentifications = append(chatIdentifications, int64(chat.(int)))
			}
		}
	} else {

		updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
			Offset:  0,
			Limit:   10,
			Timeout: 10,
		})
		wasUpdated = true
		logrus.Debug("Updates: ", updates)

		if err != nil {
			logrus.Warnf("Failed to get updates for Telegram bot with token %s: %v", token, err)
			chatIdentifications = append(chatIdentifications, 0)
		} else {
			for _, update := range updates {
				chatIdentifications = append(chatIdentifications, update.Message.Chat.ID)
			}
		}
	}

	logrus.Debug("Chat identifications: ", chatIdentifications)

	tgBot := &TelegramBot{
		ChatIdentifications: chatIdentifications,
		Client:              bot,
		Token:               textToken,
		WasUpdated:          wasUpdated,
	}

	tgBot.ChatIdentifications = chatIdentifications

	return tgBot, nil
}

func NewTransportTelegram(tokens []interface{}) Transport {
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

		if bot == nil {
			logrus.Warn("Failed to create Telegram bot with token ", token)
			continue
		}

		var textToken string
		switch token.(type) {
		case string:
			textToken = token.(string)
		case map[string]interface{}:
			textToken = token.(map[string]interface{})["token"].(string)
		}

		bots[textToken] = bot
	}

	logrus.Debug("Created Telegram transport with ", len(bots), " bots")

	transportTelegram = &TransportTelegram{Bots: bots}
	creating.Done()
	return transportTelegram
}

func (t *TransportTelegram) Deliver(block *model.Block) error {
	transactions := block.Transactions
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
