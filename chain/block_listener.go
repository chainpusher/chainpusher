package chain

import (
	"math/big"

	"github.com/chainpusher/chainpusher/model"
)

type BlockListener interface {
	BeforeQuerying(height *big.Int)

	AfterRawQuerying(block interface{}, err error)

	AfterQuerying(block *model.Block, err error)
}
