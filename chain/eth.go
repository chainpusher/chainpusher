package chain

import "github.com/chainpusher/chainpusher/sys"

const (
	InfuraApiUrl        string = "https://mainnet.infura.io/v3/"
	EthereumUsdtAddress string = "0xdac17f958d2ee523a2206206994597c13d831ec7"

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

func GetInfuraApiUrl() (string, error) {

	key, err := sys.GetEnv("INFURA_KEY")
	if err != nil {
		return key, err
	}

	return InfuraApiUrl + key, nil
}
