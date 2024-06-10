package postoffice

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/chainpusher/chainpusher/model"
	"github.com/sirupsen/logrus"
)

type TransactionCommand struct {
	Network string `json:"network"`

	CryptoCurrency string `json:"crypto_currency"`

	Payee string `json:"payee"`

	Payer string `json:"payer"`

	Amount big.Int `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
}

func ToTransactionCommand(t *model.Transaction) *TransactionCommand {
	return &TransactionCommand{
		Network:        t.Platform.String(),
		CryptoCurrency: t.CryptoCurrency.String(),
		Payee:          t.Payee,
		Payer:          t.Payer,
		Amount:         t.Amount,
		CreatedAt:      t.CreatedAt,
	}
}

type DeliveryCommand struct {
	Transactions []*TransactionCommand `json:"transactions"`
}

type TransportHttp struct {
	Urls []string
}

func (po *TransportHttp) Deliver(transactions []*model.Transaction) error {
	commands := make([]*TransactionCommand, 0)
	for _, transaction := range transactions {
		commands = append(commands, ToTransactionCommand(transaction))
	}

	payload := &DeliveryCommand{
		Transactions: commands,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logrus.Warnf("Failed to marshal payload: %v", err)
		logrus.Debugf("%v", payload)
		return err
	}

	for _, url := range po.Urls {

		go func(url string, payload []byte) {
			body := bytes.NewReader(payloadBytes)
			req, err := http.NewRequest("POST", url, body)
			if err != nil {
				logrus.Warnf("Failed to create request: %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}

			response, err := client.Do(req)
			if err != nil {
				logrus.Warnf("Failed to deliver message: %v", err)
				return
			}

			if response.StatusCode != http.StatusOK {
				logrus.Warnf("Failed to deliver message: %v", response.Status)
			}

			defer response.Body.Close()

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				logrus.Warnf("Failed to read response body: %v", err)
			}
			logrus.Debug(string(bodyBytes))

			logrus.Infof("Delivered message to %s", url)
		}(url, payloadBytes)
	}

	return nil
}

func NewTransportHttp(urls []string) Transport {
	return &TransportHttp{
		Urls: urls,
	}
}
