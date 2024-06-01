package chain

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tc "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
)

// test the log of tron contract
func TestTronUSDTContractLog(t *testing.T) {
	if "payee" != "a" {
		t.Error("payee is incorrect")
	}

	if "payer" != "b" {
		t.Error("payer is incorrect")
	}

}

func TestTronTransactionTransferContract(t *testing.T) {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	transaction, err := client.GetTransactionByID("68fc9d0cd12807e0d0ef25dba843c1e0a06f34b72e102b8ed3be051bda3c989a")
	if err != nil {
		t.Error(err)
	}

	contract := transaction.RawData.GetContract()[0]

	var transferContract core.TransferContract
	contract.Parameter.UnmarshalTo(&transferContract)

	payer := "TBREsCfBdPyD612xZnwvGPux7osbXvtzLh"
	payee := "TYqDptDgPUBwhrETDqJjjrCocKgwrQ5tyw"
	var amount int64 = 20000000

	if transferContract.Amount != amount {
		t.Error("Amount is incorrect")
	}

	if payer != tc.EncodeCheck(transferContract.OwnerAddress) {
		t.Error("Payer is incorrect")
	}

	if payee != tc.EncodeCheck(transferContract.ToAddress) {
		t.Error("Payee is incorrect")
	}
}

// test the transaction of tron
func TestTronTransaction(t *testing.T) {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	transaction, err := client.GetTransactionInfoByID("73d387d1ef336ed0ad79fdd886107b0f2942e5823db5169faf211f40958daa4d")
	if err != nil {
		t.Error(err)
	}
	addr := transaction.ContractAddress
	addrHex := hex.EncodeToString(addr)
	addrBase58 := address.Address(addr).String()

	if len(addr) != 21 {
		t.Error("Contract address length is incorrect")
	}

	if addrHex != "41a614f803b6fd780986a42c78ec9c7f77e6ded13c" {
		t.Error("Contract address is incorrect")
	}

	if addrBase58 != "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" {
		t.Error("Contract address is incorrect")
	}

	t.Log(err, addrHex)
}

// PLAN: application logic
func TestTronParseArgumentsOfSmartContract(t *testing.T) {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	input, _ := hex.DecodeString("a9059cbb00000000000000000000000011e4178f495918a287807adc40d18f53239bf91f0000000000000000000000000000000000000000000000000000000005f5e100")

	contractAbi, err := client.GetContractABI("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")

	if err != nil {
		return
	}

	args, err := abi.GetInputsParser(contractAbi, "transfer")
	if err != nil {
		t.Error(err)
	}
	unpacked, err := args.Unpack(input[4:])

	ethAddress := unpacked[0].(common.Address)
	amount := unpacked[1].(*big.Int)

	tronAddress := ToTronAddress(ethAddress)

	if err != nil {
		t.Error(err)
	}

	if tronAddress.String() != "TBbojS2CE76ury6v9zxamuzVCaDftVouN5" {
		t.Error("Address is incorrect")
	}

	t.Log(tronAddress.String(), amount)
}
