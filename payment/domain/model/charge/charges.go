package charge

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/transaction"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
)

type Charges shared.Slice[*Charge]

func (charges Charges) MatchingTransactionsHaveBeenPaid(txs shared.Slice[*transaction.Transaction]) Charges {
	return charges
}
