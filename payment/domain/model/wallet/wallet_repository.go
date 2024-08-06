package wallet

type Repository interface {
	FindWallet(applicationId int64) (*Wallet, error)
}
