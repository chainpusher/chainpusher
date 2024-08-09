package charge

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/transaction"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
)

type Repository interface {
	Find(id int64) (*Charge, error)

	FindCharging() ([]*Charge, error)

	FindChargingByTransactions(transactions shared.Slice[*transaction.Transaction]) (Charges, error)

	Complete(charges Charges) error

	Save(entity *Charge) error
}
