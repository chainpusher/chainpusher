package wallet

import "time"

type Wallet struct {
	ID         int64
	Blockchain string
	Crypto     string
	PoolId     int64
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
