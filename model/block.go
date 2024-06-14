package model

import (
	"math/big"
	"time"
)

type Block struct {
	Height *big.Int

	Id string

	Transactions []*Transaction

	CreatedAt time.Time
}

func (b *Block) GenerateTimeToNextBlock() time.Time {
	return time.Now()
}

// Mabye this block is across multiple blocks
func (b *Block) IsAcrossMultipleBlocks() bool {
	return false
}
