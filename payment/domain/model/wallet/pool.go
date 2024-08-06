package wallet

import "time"

type Pool struct {
	ID        int64
	AccountId int64
	Wallets   []Wallet
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Pool) Groups() map[string][]Wallet {
	groups := make(map[string][]Wallet)
	for _, w := range p.Wallets {
		groups[w.Blockchain] = append(groups[w.Blockchain], w)
	}
	return groups
}
