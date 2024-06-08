package chain_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/chain"
	"github.com/ethereum/go-ethereum/common"
)

func TestEthereumServiceAssemblerToTransfer(t *testing.T) {
	expectedAddress := "0x5b6d5BB6995A7C21aaC64c78A4c5B88470a0B15e"
	expectedAmount := 1

	assembler, err := chain.NewEthereumServiceAssembler()
	if err != nil {
		t.Error("Failed to create Ethereum service assembler: ", err)
		return
	}

	data := "0xa9059cbb0000000000000000000000005b6d5bb6995a7c21aac64c78a4c5b88470a0b15e0000000000000000000000000000000000000000000000000000000000000001"
	dataBytes := common.FromHex(data)
	transfer, err := assembler.ToUsdtTransferArguments(&dataBytes)

	if err != nil {
		t.Error("Failed to parse transfer arguments: ", err)
		return
	}

	if transfer.To.Hex() != expectedAddress {
		t.Error("Failed to parse transfer address")
		return
	}

	if transfer.Value.Int64() != int64(expectedAmount) {
		t.Error("Failed to parse transfer amount")
		return
	}
}
