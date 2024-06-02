package chain_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/chainpusher/chainpusher/chain"
	tcAbi "github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

func TestTronTranserContract(t *testing.T) {
	amount := big.NewInt(90000000)
	payer := "TSGXatKFbu5djAj68Wwk9mThqZN3Wht3me"
	payee := "TDqSquXBgUCLYvYC4XZgrprLK589dkhSCf"

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	transaction, err := client.GetTransactionByID("8dd8b3777f3da4d58f197e5ac69e0339c9957de8fecd27f3fb30516092de2ce8")
	if err != nil {
		t.Error(err)
	}

	contract := transaction.RawData.GetContract()[0]
	transfer := chain.FromTronTransferContract(contract)

	if transfer.Amount.Cmp(amount) != 0 {
		t.Error("Amount is incorrect")
	}

	if transfer.Payer != payer {
		t.Error("Payer is incorrect")
	}

	if transfer.Payee != payee {
		t.Error("Payee is incorrect")
	}

}

func TestTronUsdtTriggerSmartContract(t *testing.T) {
	amount := big.NewInt(390)
	payer := "TGwCHCetZXFTyvNmhY1DmU5ndTfxPDQQ1t"
	payee := "TVfDGG67P8778zZKmgaMnU1sYegfwsQaWg"

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

	transaction, err := client.GetTransactionByID("0114c243840e365b023533c8b4dec5c69c3a5585e8c30178181beb08b19e9e11")
	if err != nil {
		t.Error(err)
	}

	smart, err := client.GetContractABI(chain.TronUsdtAddress)
	if err != nil {
		t.Error(err)
	}

	args, err := tcAbi.GetInputsParser(smart, "transfer")
	if err != nil {
		t.Error(err)
	}

	contract := transaction.RawData.GetContract()[0]
	transfer, err := chain.FromTronTriggerSmartContract(&args, contract)
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
