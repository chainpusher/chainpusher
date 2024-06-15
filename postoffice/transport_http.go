package postoffice

import (
	"bytes"
	"encoding/json"
	"github.com/chainpusher/blockchain/model"
	"io"
	"math/big"
	"net/http"
	"time"

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

func (po *TransportHttp) Deliver(block *model.Block) error {
	commands := make([]*TransactionCommand, 0)
	for _, transaction := range block.Transactions {
		commands = append(commands, ToTransactionCommand(transaction))
	}

	payload := &DeliveryCommand{
		Transactions: commands,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range po.Urls {

		go func(url string, payload []byte) {
			body := bytes.NewReader(payloadBytes)
			req, err := http.NewRequest("POST", url, body)
			if err != nil {
				// TODO: unhandled error
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}

			response, err := client.Do(req)
			if err != nil {
				// TODO: unhandled error
				return
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					logrus.Warnf("Failed to close response body: %v", err)
				}
			}(response.Body)
		}(url, payloadBytes)
	}

	return nil
}

func NewTransportHttp(urls []string) Transport {
	return &TransportHttp{
		Urls: urls,
	}
}
