package application

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/transaction"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
)

type ChargeService interface {
	Charge(a *account.Account, charge *charge.Charge) (*charge.Charge, error)

	Charged(transactions shared.Slice[*transaction.Transaction]) error
}
