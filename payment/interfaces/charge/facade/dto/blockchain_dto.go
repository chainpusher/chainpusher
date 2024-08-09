package dto

import "time"

type BlockchainDTO struct {
	Id           string
	Height       int64
	Transactions []*BlockchainTransactionDTO
	CreatedAt    time.Time
}
