package charge

import "time"

type WalletPool struct {
	ID        int64
	ChargeId  int64
	Wallets   []Wallet
	CreatedAt time.Time
}
