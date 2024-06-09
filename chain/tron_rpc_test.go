package chain_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/chainpusher/chainpusher/chain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

// TestTronParseArgumentsOfSmartContract tests the parsing of the arguments of a smart contract
func TestTronParseArgumentsOfSmartContract(t *testing.T) {
	// payee := "TNEJn1gqWKbNo26TsimuazmxxpjqufdS83"
	payer := "TVfDGG67P8778zZKmgaMnU1sYegfwsQaWg"

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	input, _ := hex.DecodeString("a9059cbb000000000000000000000041d7fb3f45980187dbbe41c30fbba76f008e16e89f0000000000000000000000000000000000000000000000000000000000000186")
	if len(input) != 68 {
		t.Error("Input length is incorrect")
	}
	contractAbi, err := chain.GetUsdtSmartContract(client)

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

	tronAddress := chain.ToTronAddress(ethAddress).String()

	if err != nil {
		t.Error(err)
	}

	if tronAddress != payer {
		t.Error("Address is incorrect")
	}

	if amount.Int64() != 390 {
		t.Error("Amount is incorrect")
	}
}
