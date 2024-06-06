package postoffice_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
)

func TestTransportHttp_Deliver(t *testing.T) {
	transactions := []*model.Transaction{
		{
			Platform:       model.PlatformBigcoin,
			CryptoCurrency: model.Bitcoin,
			Payee:          "a",
			Payer:          "b",
			Amount:         *big.NewInt(100),
			CreatedAt:      time.Now(),
		},
	}

	tp := &postoffice.TransportHttp{Urls: []string{"https://httpbin.org/post"}}
	tp.Deliver(transactions)

	time.Sleep(10 * time.Second)
}
