package price

type Repository interface {
	Save(price interface{}) error
}
