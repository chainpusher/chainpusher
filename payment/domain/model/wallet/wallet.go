package wallet

import "time"

type Wallet struct {
	ID         int64
	AccountId  int64
	Blockchain string
	Crypto     string
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
