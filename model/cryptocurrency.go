package model

type CryptoCurrency int

const (
	Bitcoin CryptoCurrency = iota
	Ethereum
	EthereumUSDT
	Tron
	TronUSDT
)

func (cc CryptoCurrency) String() string {
	switch cc {
	case Bitcoin:
		return "Bitcoin"
	case Ethereum:
		return "Ethereum"
	case Tron:
		return "Tron"
	case TronUSDT:
		return "TronUSDT"
	default:
		return "Unknown"
	}
}
