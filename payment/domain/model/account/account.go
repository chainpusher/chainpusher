package account

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/secret"
	"github.com/chainpusher/chainpusher/payment/domain/model/wallet"
	"time"
)

type Account struct {
	ID        int64 `gorm:"foreignKey:AccountId"`
	Pool      wallet.Pool
	Secrets   []secret.Secret
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (account *Account) PickWallet(c *charge.Charge) *charge.WalletPool {
	// TODO: Implement this method

	groups := account.Pool.Groups()

	var pool charge.WalletPool
	var wallets []charge.Wallet
	for blockchain := range groups {
		ws := groups[blockchain]
		if len(ws) == 0 {
			continue
		}
		w := ws[0]
		wallets = append(wallets, charge.Wallet{Block: w.Blockchain, Crypto: w.Crypto, Text: w.Address})
	}
	pool.Wallets = wallets

	return &pool
}

func NewAccount() (*Account, error) {
	s, err := secret.NewSecret()
	if err != nil {
		return nil, err
	}
	a := &Account{
		Secrets:   []secret.Secret{*s},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return a, nil
}
