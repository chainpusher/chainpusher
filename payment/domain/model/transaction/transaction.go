package transaction

import "time"

type Transaction struct {
	ID     int64
	Amount int64

	Blockchain BlockchainTransaction

	CreatedAt time.Time
}
