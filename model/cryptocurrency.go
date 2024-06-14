package model

type CryptoCurrency int

const (
	Bitcoin CryptoCurrency = iota
	Ether
	EthereumUSDT
	TRX
	TronUSDT
)

func (cc CryptoCurrency) String() string {
	switch cc {
	case Bitcoin:
		return "Bitcoin"
	case Ether:
		return "Ether"
	case TRX:
		return "TRX"
	case TronUSDT:
		return "TronUSDT"
	default:
		return "Unknown"
	}
}
