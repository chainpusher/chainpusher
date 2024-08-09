package counter

type Repository interface {
	FindCounterByKey(key string) (*Counter, error)

	IncrementCounterByKey(key string) (int64, error)

	Save(counter *Counter) error
}
