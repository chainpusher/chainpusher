package service

import (
	"math/big"

	"github.com/chainpusher/blockchain/model"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

type TronBlockChainService struct {
	client    *client.GrpcClient
	assembler *TronBlockChainAssembler
	listeners []BlockListener
}

func (service *TronBlockChainService) GetLatestBlock() (*model.Block, error) {

	for _, listener := range service.listeners {
		go listener.BeforeQuerying(big.NewInt(0))
	}

	block, err := service.client.GetNowBlock()
	if err != nil {
		return nil, err
	}
	for _, listener := range service.listeners {
		go listener.AfterRawQuerying(block, nil)
	}

	aBlock, err := service.assembler.ToBlock(block, service)
	for _, listener := range service.listeners {
		go listener.AfterQuerying(aBlock, err)
	}

	return aBlock, err
}

func (service *TronBlockChainService) GetBlock(height *big.Int) (*model.Block, error) {
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

func (service *TronBlockChainService) GetSmartContractABI(address string) (*core.SmartContract_ABI, error) {
	return service.client.GetContractABI(address)
}

func NewTronBlockChainService(client *client.GrpcClient, listeners []BlockListener) *TronBlockChainService {
	return &TronBlockChainService{
		client:    client,
		assembler: NewTronBlockChainAssembler(),
		listeners: listeners,
	}
}
