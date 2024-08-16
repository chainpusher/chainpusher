package exchange

type Service interface {
	GetPrice(cryptos ...string) ([]*Price, error)
}
