package chain

import (
	"math/big"

	"github.com/chainpusher/chainpusher/model"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

type TronV2BlockChainService struct {
	client    *client.GrpcClient
	assembler *TronBlockChainAssembler
	listeners []BlockListener
}

func (service *TronV2BlockChainService) GetLatestBlock() (*model.Block, error) {

	for _, listener := range service.listeners {
		listener.BeforeQuerying(big.NewInt(0))
	}

	block, err := service.client.GetNowBlock()
	if err != nil {
		return nil, err
	}
	for _, listener := range service.listeners {
		listener.AfterRawQuerying(block, nil)
	}

	aBlock, err := service.assembler.ToBlock(block, service)
	for _, listener := range service.listeners {
		listener.AfterQuerying(aBlock, err)
	}

	return aBlock, err
}

func (service *TronV2BlockChainService) GetBlock(height *big.Int) (*model.Block, error) {
	for _, listener := range service.listeners {
		listener.BeforeQuerying(height)
	}

	block, err := service.client.GetBlockByNum(height.Int64())
	if err != nil {
		return nil, err
	}
	for _, listener := range service.listeners {
		listener.AfterRawQuerying(block, nil)
	}

	aBlock, err := service.assembler.ToBlock(block, service)
	for _, listener := range service.listeners {
		listener.AfterQuerying(aBlock, err)
	}

	return aBlock, err
}

func (service *TronV2BlockChainService) GetSmartContractABI(address string) (*core.SmartContract_ABI, error) {
	return service.client.GetContractABI(address)
}

func NewTronV2BlockChainService() *TronV2BlockChainService {
	return &TronV2BlockChainService{}
}
