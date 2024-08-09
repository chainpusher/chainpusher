package counter

type Repository interface {
	FindCounterByKey(key string) (*Counter, error)

	Save(counter *Counter) error
}
