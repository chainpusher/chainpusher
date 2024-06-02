package model

type CryptoCurrency int

const (
	Bitcoin CryptoCurrency = iota
	Etherium
	Tron
	TronUSDT
)

func (cc CryptoCurrency) String() string {
	switch cc {
	case Bitcoin:
		return "Bitcoin"
	case Etherium:
		return "Etherium"
	case Tron:
		return "Tron"
	case TronUSDT:
		return "TronUSDT"
	default:
		return "Unknown"
	}
}
