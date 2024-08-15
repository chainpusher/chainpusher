package price

import "time"

type Price struct {
	Id         int64
	Blockchain string
	Crypto     string
	Price      int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
