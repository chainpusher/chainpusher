package transaction

import "time"

type Transaction struct {
	ID                      int64
	ChargeId                int64
	Amount                  int64
	Blockchain              Blockchain
	BlockchainTransactionId string
	Payer                   string
	Payee                   string
	CreatedAt               time.Time
}
