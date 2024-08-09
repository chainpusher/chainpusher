package counter

import "time"

type Counter struct {
	Id        int64
	Key       string
	Value     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
