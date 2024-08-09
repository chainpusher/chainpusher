package dto

type BlockchainTransactionDTO struct {
	Id     string
	Payer  string
	Payee  string
	Amount int64
}
