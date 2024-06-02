package model

type Platform int

const (
	PlatformBigcoin Platform = iota
	PlatformEtherium
	PlatformTron
)

func (p Platform) String() string {
	switch p {
	case PlatformBigcoin:
		return "Bigcoin"
	case PlatformEtherium:
		return "Etherium"
	case PlatformTron:
		return "Tron"
	default:
		return "Unknown"
	}
}
