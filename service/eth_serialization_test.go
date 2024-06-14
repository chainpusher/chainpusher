package service_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/chainpusher/blockchain/service"
	"github.com/stretchr/testify/assert"
)

// Test the transaction serialization and deserialization
func TestEthereumTransactionsSerialize(t *testing.T) {
	url, err := service.GetInfuraApiUrlFromEnvironmentVariable()
	if err != nil {
		return
	}

	svc, err := service.NewEthereumClient(url, nil)
	assert.Nil(t, err, "Failed to create Ethereum block chain service: ", err)

	block, err := svc.Client.BlockByNumber(context.Background(), big.NewInt(20045824))
	assert.Nilf(t, err, "Failed to get block: %v", err)

	transactions := block.Transactions()

	var txs service.EthereumTransactions = make(service.EthereumTransactions, len(transactions))
	txs.FromTransactions(block, transactions)

	jsonBytes, err := json.Marshal(txs)

	assert.Nilf(t, err, "Failed to marshal transactions: %v", err)
	assert.Truef(t, len(jsonBytes) > 0, "Failed to marshal transactions: empty")
}
