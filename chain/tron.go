package chain

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
)

type ContractParameterValueProtoBuf struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`

	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,proto3" json:"owner_address,omitempty"`

	ContractAddress []byte `protobuf:"bytes,3,opt,name=contract_address,proto3" json:"contract_address,omitempty"`
}

type ContractParameterValueDTO struct {
	Data string `json:"data"`

	OwnerAddress string `json:"owner_address"`

	ContractAddress string `json:"contract_address"`
}

func GetTronContractParamter(contract *core.Transaction_Contract) *ContractParameter {

	param := &ContractParameter{}
	proto.Unmarshal(contract.Parameter.Value, param)

	return param
}

func GetSmartCratractDTO(client *client.GrpcClient, address address.Address) (*core.SmartContract_ABI, error) {
	abi, err := client.GetContractABI(address.String())
	if err != nil {
		return nil, err
	}
	return abi, nil
}
