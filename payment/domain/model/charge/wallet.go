package charge

type Wallet struct {
	ID         int64
	ChargeId   int64
	Blockchain string
	Crypto     string
	Address    string
}

func (Wallet) TableName() string {
	return "charge_wallets"
}
