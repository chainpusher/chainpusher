package account

import (
	"fmt"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/secret"
	"github.com/chainpusher/chainpusher/payment/domain/model/wallet"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"time"
)

type Account struct {
	ID        int64 `gorm:"foreignKey:AccountId"`
	Wallets   []wallet.Wallet
	Secrets   []secret.Secret
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (account *Account) PickWallets() shared.Slice[charge.Wallet] {
	wallets := make(shared.Slice[charge.Wallet], 0)
	grouped := shared.GroupBy(shared.Of(account.Wallets...), func(w wallet.Wallet) string {
		return fmt.Sprintf("%s-%s", w.Blockchain, w.Crypto)
	})

	grouped.ForEach(func(k string, v shared.Slice[wallet.Wallet]) {
		w := charge.Wallet{
			Blockchain: v[0].Blockchain,
			Crypto:     v[0].Crypto,
			Text:       v[0].Address,
		}
		wallets = append(wallets, w)
	})

	return wallets
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
