package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
)

// ToTronAddress converts an Ethereum address to a Tron address.
func ToTronAddress(addr common.Address) address.Address {
	var tronAddress = make([]byte, len(addr)+1)
	tronAddress[0] = 0x41
	copy(tronAddress[1:], addr.Bytes())

	return address.Address(tronAddress)
}

func GetTronContractParamter(contract *core.Transaction_Contract) *ContractParameter {

	param := &ContractParameter{}
	proto.Unmarshal(contract.Parameter.Value, param)

	return param
}
