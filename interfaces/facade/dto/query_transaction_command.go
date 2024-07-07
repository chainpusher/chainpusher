package dto

import "time"

type QueryTransactionsCommand struct {
	BlockId int64

	Page int

	PageSize int

	Since time.Time
}
