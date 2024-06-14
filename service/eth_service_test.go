package service_test

import (
	"math/big"
	"testing"

	"github.com/chainpusher/blockchain/service"
	"github.com/stretchr/testify/assert"
)

func TestEthereumService_GetLatestBlock(t *testing.T) {
	url, err := service.GetInfuraApiUrlFromEnvironmentVariable()
	if err != nil {
		return
	}
	s, err := service.NewEthereumBlockChainService(url, []service.BlockListener{})
	assert.Nil(t, err)

	s.GetLatestBlock()
}

func TestEthereumService_GetBlock(t *testing.T) {
	url, err := service.GetInfuraApiUrlFromEnvironmentVariable()
	if err != nil {
		return
	}
	s, err := service.NewEthereumBlockChainService(url, []service.BlockListener{})
	var svc service.BlockchainService = s
	assert.Nil(t, err)

	height := big.NewInt(20045182)
	block, err := svc.GetBlock(height)
	assert.Nil(t, err)

	assert.Equal(t, height, block.Height, "Block height should be the same")
}
