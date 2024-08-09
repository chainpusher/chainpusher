package price

import "time"

type Price struct {
	Id        int64
	Amount    int64
	State     int
	Used      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
