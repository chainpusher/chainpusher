package chain

import (
	"bytes"
	"encoding/json"
	"log"
	"math/big"
	"strings"

	"github.com/chainpusher/chainpusher/infrastructure"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

type TronContract struct {
	Abi *abi.ABI
}

type TronAbiParameter []byte

func (p TronAbiParameter) Get() []byte {
	return p // p[4:]
}

func (c *TronContract) Unpack(parameter TronAbiParameter, methodName string) ([]interface{}, error) {
	bytes := parameter.Get()
	return c.Abi.Methods[methodName].Inputs.Unpack(bytes)
}

func (c *TronContract) UnpackTransfer(parameter TronAbiParameter) (address.Address, *big.Int, error) {
	return address.Address{}, nil, nil
}

// FixAbi fixes the ABI JSON, converting the type to lowercase.
func FixAbi(contract *GetContractResponse) {
	entrys := contract.Abi.Entrys
	for _, entry := range entrys {

		e := entry.(map[string]interface{})
		e["type"] = strings.ToLower(e["type"].(string))
	}
}

// TODO:
func (c *TronContract) GetContractAbi(address string) {

}

func NewTronContract(contractAddress string) (*TronContract, error) {

	// contracts := transaction.RawData.GetContract()
	// if len(contracts) == 0 {
	// 	return nil, errors.New("no contract found")
	// }

	// contract := contracts[0]
	// parameter := GetTronContractParamter(contract)
	// contractAddress := common.BytesToHexString(parameter.ContractAddress)

	contract := &TronContract{
		Abi: nil,
	}

	abiJson, err := infrastructure.GetKey(contractAddress)
	if err != nil {

		tronClient := NewTronHttpClient()
		contractDto, err := tronClient.GetContract(contractAddress)
		FixAbi(contractDto)

		if err != nil {
			return nil, err
		}

		abiJson, err = json.Marshal(contractDto.Abi.Entrys)
		if err != nil {
			return nil, err
		}

		err = infrastructure.SetKey(contractAddress, abiJson)
		if err != nil {
			return nil, err
		}
	}

	abi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return nil, err
	}

	log.Println(abi)

	return contract, nil
}
