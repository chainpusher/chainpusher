package wallet

import "time"

type CryptoWallet struct {
	ID        int64
	Crypto    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
