package service

import (
	"math/big"

	"github.com/chainpusher/blockchain/model"
)

type BlockListener interface {
	BeforeQuerying(height *big.Int)

	AfterRawQuerying(block interface{}, err error)

	AfterQuerying(block *model.Block, err error)
}
