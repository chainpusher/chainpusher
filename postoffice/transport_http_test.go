package postoffice_test

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"

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
	block := &model.Block{Transactions: transactions}

	tp := &postoffice.TransportHttp{Urls: []string{"https://httpbin.org/post"}}
	err := tp.Deliver(block)
	assert.Errorf(t, err, "Post https://httpbin.org/post: dial tcp: lookup httpbin.org: no such host")

	time.Sleep(10 * time.Second)
}
