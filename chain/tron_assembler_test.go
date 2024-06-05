package chain_test

import (
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/chain"
	tcAbi "github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"google.golang.org/grpc"
)

func TestTronTranserContract(t *testing.T) {
	amount := big.NewInt(90000000)
	payer := "TSGXatKFbu5djAj68Wwk9mThqZN3Wht3me"
	payee := "TDqSquXBgUCLYvYC4XZgrprLK589dkhSCf"

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	block, err := client.GetBlockByNum(61793343)
	if err != nil {
		t.Error(err)
	}
	var transaction *api.TransactionExtention
	for _, item := range block.Transactions {
		if hex.EncodeToString(item.Txid) == "8dd8b3777f3da4d58f197e5ac69e0339c9957de8fecd27f3fb30516092de2ce8" {
			transaction = item
		}
	}

	contract := transaction.Transaction.RawData.GetContract()[0]
	transfer := chain.FromTronTransferContract(transaction, contract)

	if transfer.Amount.Cmp(amount) != 0 {
		t.Error("Amount is incorrect")
	}

	if transfer.Payer != payer {
		t.Error("Payer is incorrect")
	}

	if transfer.Payee != payee {
		t.Error("Payee is incorrect")
	}

	if transfer.CreatedAt.Unix() != 1716023160000 {
		t.Error("CreatedAt is incorrect")

	}

}

func TestTronUsdtTriggerSmartContract(t *testing.T) {
	amount := big.NewInt(5000000000)
	payer := "TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R"
	payee := "TL8TBpubVzBr1UWPXBXU8Pci5ZAip9SwEf"

	payee58, err := address.Base58ToAddress(payee)
	if err != nil {
		t.Error(err)
	}
	t.Log("payee hex", hex.EncodeToString(payee58.Bytes()))

	payer58, err := address.Base58ToAddress(payer)
	if err != nil {
		t.Error(err)
	}
	t.Log("payer hex", hex.EncodeToString(payer58.Bytes()))

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())
	client.SetTimeout(60 * time.Second)

	block, err := client.GetBlockByNum(61793343)
	if err != nil {
		t.Error(err)
	}
	var transaction *api.TransactionExtention
	for _, item := range block.Transactions {
		if hex.EncodeToString(item.Txid) == "fd09b652145ef77a3d8b09b235bd601c344a56f46b0fe56d590f9d967716ea5d" {
			transaction = item
		}
	}

	smart, err := client.GetContractABI(chain.TronUsdtAddress)
	if err != nil {
		t.Error(err)
	}

	args, err := tcAbi.GetInputsParser(smart, "transfer")
	if err != nil {
		t.Error(err)
	}

	contract := transaction.Transaction.RawData.GetContract()[0]
	transfer, err := chain.FromTronTriggerSmartContract(&args, transaction, contract)
	if err != nil {
		t.Error(err)
	}

	if transfer.Amount.Cmp(amount) != 0 {
		t.Error("Amount is incorrect", transfer.Amount)
	}

	if transfer.Payer != payer {
		t.Error("Payer is incorrect", transfer.Payer)
	}

	if transfer.Payee != payee {
		t.Error("Payee is incorrect", transfer.Payee)
	}

	if transfer.CreatedAt.Unix() != 1716023162778 {
		t.Error("CreatedAt is incorrect", transfer.CreatedAt.Unix())
	}

}

func TestGetUsdtSmartContract(t *testing.T) {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	abi, err := chain.GetUsdtSmartContract(client)
	if err != nil {
		t.Error(err)
	}

	if abi == nil {
		t.Error("ABI is nil")
	}
}
