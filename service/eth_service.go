package service

import (
	"context"
	"math/big"

	"github.com/chainpusher/blockchain/model"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumBlockChainService struct {
	client    *ethclient.Client
	assembler *EthereumServiceAssembler
	listeners []BlockListener
}

func (s *EthereumBlockChainService) GetLatestBlock() (*model.Block, error) {
	for _, listener := range s.listeners {
		go listener.BeforeQuerying(big.NewInt(0))
	}

	block, err := s.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	for _, listener := range s.listeners {
		go listener.AfterRawQuerying(block, nil)
	}

	aBlock := s.assembler.ToBlock(block)
	for _, listener := range s.listeners {
		go listener.AfterQuerying(aBlock, err)
	}

	return aBlock, nil
}

func (s *EthereumBlockChainService) GetBlock(height *big.Int) (*model.Block, error) {
	for _, listener := range s.listeners {
		go listener.BeforeQuerying(big.NewInt(0))
	}

	block, err := s.client.BlockByNumber(context.Background(), height)
	if err != nil {
		return nil, err
	}
	for _, listener := range s.listeners {
		go listener.AfterRawQuerying(block, nil)
	}

	aBlock := s.assembler.ToBlock(block)
	for _, listener := range s.listeners {
		go listener.AfterQuerying(aBlock, err)
	}

	return aBlock, nil
}

func NewEthereumBlockChainService(url string, listeners []BlockListener) (*EthereumBlockChainService, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	assembler, err := NewEthereumServiceAssembler()
	if err != nil {
		return nil, err
	}

	return &EthereumBlockChainService{
		client:    client,
		assembler: assembler,
		listeners: listeners,
	}, nil
}
