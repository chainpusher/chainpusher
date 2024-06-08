package postoffice_test

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
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

	tg := postoffice.NewTransportTelegram([]string{token}).(*postoffice.TransportTelegram)

	if err != nil {
		t.Error("Failed to create Telegram transport: ", err)
	}

	tg.Deliver([]*model.Transaction{
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ethereum,
			Payee:          "payee",
			Payer:          "payer",
			Amount:         *big.NewInt(100),
		},
		{
			Platform:       model.PlatformEthereum,
			CryptoCurrency: model.Ethereum,
			Payee:          "payee2",
			Payer:          "payer2",
			Amount:         *big.NewInt(200),
		},
	})

	t.Log("Telegram transport created successfully: ", tg)
	time.Sleep(10 * time.Second)
}
