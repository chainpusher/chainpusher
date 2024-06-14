package service

import (
	"context"
	"math/big"

	"github.com/chainpusher/blockchain/model"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	Client    *ethclient.Client
	Assembler *EthereumServiceAssembler
	Channel   chan interface{}
}

func (s *EthereumClient) GetNowBlock() (*types.Header, error) {

	header, err := s.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func (s *EthereumClient) GetBlock(number *big.Int) (*types.Block, error) {

	var block *types.Block

	block, err := s.Client.BlockByNumber(context.Background(), number)

	go func(block *types.Block) {
		if s.Channel == nil {
			return
		}

		var txs EthereumTransactions = make(EthereumTransactions, len(block.Transactions()))
		txs.FromTransactions(block, block.Transactions())

		s.Channel <- txs
	}(block)

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *EthereumClient) GetTransactions(number *big.Int) ([]*model.Transaction, error) {

	block, err := s.GetBlock(number)
	if err != nil {
		return nil, err
	}

	return s.Assembler.BlockToTransactions(block), nil
}

func NewEthereumClient(url string, channel chan interface{}) (*EthereumClient, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	assembler, err := NewEthereumServiceAssembler()
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		Client:    client,
		Assembler: assembler,
		Channel:   channel,
	}, nil
}
