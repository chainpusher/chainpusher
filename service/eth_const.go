package service

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const (
	InfuraApiUrl         string = "https://mainnet.infura.io/v3/"
	EthereumUsdtAddress  string = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	EthereumEmptyAddress string = "0x0000000000000000000000000000000000000000"

	EthereumUsdtAbi string = `[
		{
			"constant": false,
			"inputs": [
				{
					"name": "_to",
					"type": "address"
				},
				{
					"name": "_value",
					"type": "uint256"
				}
			],
			"name": "transfer",
			"outputs": [
				{
					"name": "",
					"type": "bool"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`
)

var EthereumUsdtMethodTransfer []byte = []byte{0xa9, 0x5, 0x9c, 0xbb}

type EthereumContractUsdtTransfer struct {
	To    *common.Address
	Value *big.Int
}

func GetInfuraApiUrlFromEnvironmentVariable() (string, error) {

	return InfuraApiUrl + os.Getenv("INFURA_KEY"), nil
}

func GetInfuraApiUrlV2(key string) string {
	return InfuraApiUrl + key
}
