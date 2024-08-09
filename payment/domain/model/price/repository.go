package price

type Repository interface {
	FindPriceByAmount(amount int64) (*Price, error)

	Save(price *Price) error
}
