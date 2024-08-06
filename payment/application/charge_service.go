package application

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
)

type ChargeService interface {
	Charge(a *account.Account, charge *charge.Charge) (*charge.Charge, error)

	Charged(c *charge.Charge) error
}
