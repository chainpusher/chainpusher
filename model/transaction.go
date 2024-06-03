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

func (t *Transaction) SocialMediaMessage() string {
	return fmt.Sprintf("%s address recived %s %s from %s on %s", t.Payee, t.Amount.String(), t.CryptoCurrency.String(), t.Payer, t.Platform.String())
}

// func GetAddresses(transactions []*Transaction) []string {
// 	set := make(map[string]int)
// 	for _, transaction := range transactions {
// 		set[transaction.Payee] = 1
// 		set[transaction.Payer] = 1
// 	}

// 	addresses := make([]string, 0)
// 	for address := range set {
// 		addresses := append(addresses, address)
// 	}
// 	return addresses
// }
