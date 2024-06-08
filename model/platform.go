package model

type Platform int

const (
	PlatformBigcoin Platform = iota
	PlatformEthereum
	PlatformTron
)

func (p Platform) String() string {
	switch p {
	case PlatformBigcoin:
		return "Bigcoin"
	case PlatformEthereum:
		return "Ethereum"
	case PlatformTron:
		return "Tron"
	default:
		return "Unknown"
	}
}
