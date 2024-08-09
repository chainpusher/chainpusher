package counter

import "time"

type Counter struct {
	Id        int64
	Key       string
	Value     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCounter(key string) *Counter {
	return &Counter{
		Key:       key,
		Value:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
