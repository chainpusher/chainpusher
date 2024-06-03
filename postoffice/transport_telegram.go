package postoffice

import (
	"strings"

	"github.com/chainpusher/chainpusher/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TelegramBot struct {
	ChatIdentifications []int64

	Client *tgbotapi.BotAPI
}

type TransportTelegram struct {
	Bots map[string]*TelegramBot
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Warnf("Failed to create Telegram bot with token %s: %v", token, err)
		return nil, err
	}

	logrus.Debug("Authorized on account ", bot.Self.UserName)
	updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   10,
		Timeout: 60,
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

	bots := make(map[string]*TelegramBot)

	for _, token := range tokens {
		bot, err := NewTelegramBot(token)
		if err != nil {
			continue
		}

		bots[token] = bot
	}

	return &TransportTelegram{Bots: bots}
}

func (t *TransportTelegram) Deliver(transactions []*model.Transaction) error {
	messages := make([]string, 0)
	for _, transaction := range transactions {
		messages = append(messages, transaction.SocialMediaMessage())
	}
	message := strings.Join(messages, "\n")

	for _, bot := range t.Bots {
		go func(bot *TelegramBot, message string) {
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
