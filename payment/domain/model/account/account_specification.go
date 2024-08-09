package account

type AccountSpecification struct {
	secretKey string

	walletAddress string
}

func (a *AccountSpecification) SecretKey() string {
	return a.secretKey
}

func (a *AccountSpecification) WalletAddress() string {
	return a.walletAddress
}

func NewAccountSpecification(secretKey string, walletAddress string) *AccountSpecification {
	return &AccountSpecification{
		secretKey:     secretKey,
		walletAddress: walletAddress,
	}
}
