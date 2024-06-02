package model

import (
	"fmt"
	"math/big"
)

type Transaction struct {
	Platform Platform

	CryptoCurrency CryptoCurrency

	Payee string

	Payer string

	Amount big.Int
}

func (t *Transaction) Logging() string {
	return fmt.Sprintf("Platform: %s, CryptoCurrency: %s, Payee: %s, Payer: %s, Amount: %s", t.Platform, t.CryptoCurrency, t.Payee, t.Payer, t.Amount.String())
}
