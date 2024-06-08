package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/chainpusher/chainpusher/infrastructure"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

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

	CacheBlockByNumber = "ethereum_block_%s"
)

func GetInfuraApiUrl() (string, error) {

	key, err := sys.GetEnv("INFURA_KEY")
	if err != nil {
		return key, err
	}

	return InfuraApiUrl + key, nil
}

type EthereumBlockChainService struct {
	Client *ethclient.Client
}

func (s *EthereumBlockChainService) GetNowBlock() (*types.Header, error) {

	header, err := s.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func (s *EthereumBlockChainService) GetBlock(number *big.Int) (*types.Block, error) {

	cacheBlockKey := fmt.Sprintf(CacheBlockByNumber, number.String())
	cacheBlock, err := infrastructure.GetKey(cacheBlockKey)
	var block *types.Block

	if err == nil {
		if err := json.Unmarshal(cacheBlock, &block); err == nil {
			return block, nil
		}
	}

	block, err = s.Client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}
	cacheBlock, err = json.Marshal(block)
	if err != nil {
		logrus.Warnf("Failed to cache block: %v", err)
	}
	infrastructure.SetKey(cacheBlockKey, cacheBlock)

	return block, nil
}

func NewEthereumBlockChainService(url string) (*EthereumBlockChainService, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &EthereumBlockChainService{Client: client}, nil
}
