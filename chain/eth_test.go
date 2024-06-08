package chain_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTronBlockChainService_GetNowBlock(t *testing.T) {
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
	header, err := service.GetNowBlock()
	if err != nil {
		t.Fatal("Failed to get header: ", err)
		return
	}

	if header.Number.Cmp(big.NewInt(1)) == -1 {
		t.Fatal("Block number is less than 1")
	}

	block, err := service.GetBlock(header.Number)

	if err != nil {
		t.Fatal("Failed to get block: ", err)
		return
	}

	b, err := service.GetBlock(big.NewInt(20045182))
	if err != nil {
		t.Fatal("Failed to get block: ", err)
		return

	}

	t.Log(block.Transactions().Len(), b)
}

func TestEthereumGetTransactionValueGreaterThan(t *testing.T) {
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

	block, err := service.GetBlock(big.NewInt(20045182))
	if err != nil {
		t.Fatal("Failed to get block: ", err)
		return

	}

	for _, tx := range block.Transactions() {
		if tx.Value().Cmp(big.NewInt(0)) == 1 {
			t.Log(tx.Value())
		}
	}
	t.Log("Transactions: ", block.Transactions().Len())
	t.Log("Current time: ", time.Now())
}

func TestEthereumGetTransactionOfUSDT(t *testing.T) {
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

	block, err := service.GetBlock(big.NewInt(20045824))
	if err != nil {
		t.Fatal("Failed to get block: ", err)
		return

	}

	for _, tx := range block.Transactions() {

		if tx.To().String() == chain.EthereumUsdtAddress {
			// t.Logf("%s and %s", tx.To(), chain.EthereumUsdtAddress)
			t.Log(tx.Value())
		}
	}
	t.Log("Transactions: ", block.Transactions().Len())
	t.Log("Current time: ", time.Now())
}

func TestEthereumGetTransactions(t *testing.T) {
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

	transactions, err := service.GetTransactions(big.NewInt(20045824))

	if err != nil {
		t.Fatal("Failed to get transactions: ", err)
		return
	}

	t.Log("Transactions: ", len(transactions))
	t.Log("Current time: ", time.Now())
}

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

	if header.Number.Cmp(big.NewInt(1)) == -1 {
		t.Fatal("Block number is less than 1")
	}
}

func TestEthereumAbiParse(t *testing.T) {
	expectedAddress := "0x5b6d5BB6995A7C21aaC64c78A4c5B88470a0B15e"
	expectedAmount := 1

	parsedABI, err := abi.JSON(strings.NewReader(chain.EthereumUsdtAbi))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	data := "0xa9059cbb0000000000000000000000005b6d5bb6995a7c21aac64c78a4c5b88470a0b15e0000000000000000000000000000000000000000000000000000000000000001" // Replace with your data
	method, err := parsedABI.MethodById(common.FromHex(data))
	if err != nil {
		log.Fatalf("Failed to get method: %v", err)
	}

	args, err := method.Inputs.Unpack(common.FromHex(data)[4:])
	if err != nil {
		log.Fatalf("Failed to unpack data: %v", err)
	}

	to := args[0].(common.Address)
	value := args[1].(*big.Int)

	if to.Hex() != expectedAddress {
		t.Fatalf("Expected address: %s, got: %s", expectedAddress, to.Hex())
	}

	if value.Int64() != int64(expectedAmount) {
		t.Fatalf("Expected amount: %d, got: %d", expectedAmount, value.Int64())
	}
}

func TestEthereumDataMethod(t *testing.T) {
	data := "0xa9059cbb0000000000000000000000005b6d5bb6995a7c21aac64c78a4c5b88470a0b15e0000000000000000000000000000000000000000000000000000000000000001"
	dataBytes := common.FromHex(data)

	if !bytes.Equal(dataBytes[:4], chain.EthereumUsdtMethodTransfer) {
		t.Fatalf("Expected method: %v, got: %v", chain.EthereumUsdtMethodTransfer, dataBytes[:4])
	}
}

func TestEthereumSerialize(t *testing.T) {
	url, err := chain.GetInfuraApiUrl()
	if err != nil {
		t.Log("Failed to get Infura API URL: ", err)
		return
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Fatal("Failed to dial Infura API: ", err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal("Failed to get header: ", err)
	}

	_, err = json.Marshal(header)
	if err != nil {
		t.Fatal("Failed to marshal header: ", err)
	}
}
