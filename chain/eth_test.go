package chain_test

import (
	"context"
	"testing"

	"github.com/chainpusher/chainpusher/sys"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestEthBlockChain_GetNowBlock(t *testing.T) {

	key, err := sys.GetEnv("INFURA_KEY")

	if err != nil {
		t.Log("Failed to get Infura key: ", err)
		return
	}

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + key)
	if err != nil {
		t.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(header)
}
