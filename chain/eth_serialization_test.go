package chain_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/chain"
)

// Test the transaction serialization and deserialization
func TestEthereumTransactionsSerialize(t *testing.T) {
	t.Log("Current time: ", time.Now())
	url, err := chain.GetInfuraApiUrl()
	if err != nil {
		t.Log("Failed to get Tron API URL: ", err)
		return
	}
	service, err := chain.NewEthereumBlockChainService(url)
	if err != nil {
		t.Fatal("Failed to create Ethereum block chain service: ", err)
		return
	}

	block, err := service.Client.BlockByNumber(context.Background(), big.NewInt(20045824))
	if err != nil {
		t.Fatal("Failed to get transactions: ", err)
		return
	}

	transactions := block.Transactions()

	var txs chain.EthereumTransactions = make(chain.EthereumTransactions, len(transactions))

	txs.FromTransactions(block, transactions)

	var o interface{} = txs
	jsonBytes, err := json.Marshal(o)

	if err != nil {
		t.Fatal("Failed to marshal transactions: ", err)
		return
	}
	t.Log(string(jsonBytes))
	t.Log("Transactions: ", len(txs))
	t.Log("Transactions: ", len(transactions))
	t.Log("Current time: ", time.Now())
}
