package chain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/chainpusher/chainpusher/infrastructure"
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

const (
	InfuraApiUrl        string = "https://mainnet.infura.io/v3/"
	EthereumUsdtAddress string = "0xdAC17F958D2ee523a2206206994597C13D831ec7"

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

var EthereumUsdtMethodTransfer []byte = []byte{0xa9, 0x5, 0x9c, 0xbb}

type EthereumContractUsdtTransfer struct {
	To    *common.Address
	Value *big.Int
}

func GetInfuraApiUrl() (string, error) {

	key, err := sys.GetEnv("INFURA_KEY")
	if err != nil {
		return key, err
	}

	return InfuraApiUrl + key, nil
}

func GetInfuraApiUrlV2(key string) string {
	return InfuraApiUrl + key
}

type EthereumBlockChainService struct {
	Client    *ethclient.Client
	Assembler *EthereumServiceAssembler
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
	err = errors.New("clear the cache")
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

func (s *EthereumBlockChainService) GetTransactions(number *big.Int) ([]*model.Transaction, error) {

	block, err := s.GetBlock(number)
	if err != nil {
		return nil, err
	}

	return s.Assembler.BlockToTransactions(block), nil
}

func NewEthereumBlockChainService(url string) (*EthereumBlockChainService, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	assembler, err := NewEthereumServiceAssembler()
	if err != nil {
		return nil, err
	}

	return &EthereumBlockChainService{
		Client:    client,
		Assembler: assembler,
	}, nil
}
