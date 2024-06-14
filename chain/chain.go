package chain

import (
	"math/big"

	"github.com/chainpusher/chainpusher/model"
)

func GetAllPlatform() []model.Platform {
	return []model.Platform{
		model.PlatformTron,
		// model.PlatformEthereum,
	}
}

type BlockchainService interface {
	GetLatestBlock() (*model.Block, error)

	GetBlock(height *big.Int) (*model.Block, error)
}
