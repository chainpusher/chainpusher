package postoffice_test

import (
	"github.com/chainpusher/blockchain/model"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/postoffice"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestNewTransportTelegram(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	cwd, err := os.Getwd()
	if err != nil {
		logrus.Error("Failed to get current working directory: ", err)
	}
	dotenvFilePath := filepath.Join(cwd, "..", ".env")
	godotenv.Load(dotenvFilePath)

	logrus.Info("Current working directory: ", cwd)

	token := os.Getenv("TEST_TELEGRAM_TOKEN")
	if len(token) == 0 {
		return
	}

	tokens := []string{token}
	anyTokens := make([]interface{}, len(tokens))
	anyTokens[0] = tokens[0]

	tg := postoffice.NewTransportTelegram(anyTokens).(*postoffice.TransportTelegram)
	block := &model.Block{
		Transactions: []*model.Transaction{
			{
				Platform:       model.PlatformEthereum,
				CryptoCurrency: model.Ether,
				Payee:          "payee",
				Payer:          "payer",
				Amount:         *big.NewInt(100),
			},
			{
				Platform:       model.PlatformEthereum,
				CryptoCurrency: model.Ether,
				Payee:          "payee2",
				Payer:          "payer2",
				Amount:         *big.NewInt(200),
			},
		},
	}
	tg.Deliver(block)

	t.Log("Telegram transport created successfully: ", tg)
	time.Sleep(5 * time.Second)
}

// test when token is an object
func TestNewTransportTelegramObject(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	e, _ := sys.GetEnv("TEST_TELEGRAM_TOKEN")
	if e == "" {
		t.Skip("Telegram token is not set")
	}

	tokens := []interface{}{
		map[string]interface{}{
			"token":   e,
			"chat_id": []int{641234},
		},
	}

	token := tokens[0].(map[string]interface{})

	bot, err := postoffice.NewTelegramBot(token)

	if err != nil {
		t.Error("Failed to create Telegram bot: ", err)
	}

	if bot == nil {
		t.Error("Expected bot")
	}

	t.Log("Telegram bot created successfully")

	if token["token"].(string) != e {
		t.Error("Expected ", e)
	}

	if token["chat_id"].([]int)[0] != 641234 {
		t.Error("Expected 641234")
	}

}

// test when token is an object but chat_id is empty
func TestNewTransportTelegramObjectChatIdEmpty(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	e, _ := sys.GetEnv("TEST_TELEGRAM_TOKEN")
	if e == "" {
		t.Skip("Telegram token is not set")
	}

	tokens := []interface{}{
		map[string]interface{}{
			"token": e,
		},
	}

	token := tokens[0].(map[string]interface{})

	bot, err := postoffice.NewTelegramBot(token)

	if err != nil {
		t.Error("Failed to create Telegram bot: ", err)
	}

	t.Log("Telegram bot created successfully")

	if bot == nil {
		t.Error("Expected bot")
		return
	}

	if bot.WasUpdated == false {
		t.Error("Expected true")
	}
}
