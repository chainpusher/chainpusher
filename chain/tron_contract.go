package chain

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/chainpusher/chainpusher/infrastructure"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type TronContract struct {
	Abi *abi.ABI
}

func FixAbi(contract *GetContractResponse) {
	entrys := contract.Abi.Entrys
	for _, entry := range entrys {
		// TODO: Fix the ABI
	}
}

func NewTronContract(contractAddress string) (*TronContract, error) {

	// contracts := transaction.RawData.GetContract()
	// if len(contracts) == 0 {
	// 	return nil, errors.New("no contract found")
	// }

	// contract := contracts[0]
	// parameter := GetTronContractParamter(contract)
	// contractAddress := common.BytesToHexString(parameter.ContractAddress)

	abiJson, err := infrastructure.GetKey(contractAddress)
	if err != nil {

		tronClient := NewTronHttpClient()
		contractDto, err := tronClient.GetContract(contractAddress)

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

	log.Println("ABI JSON: ", string(abiJson))

	abi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return nil, err
	}

	return &TronContract{
		Abi: &abi,
	}, nil
}
