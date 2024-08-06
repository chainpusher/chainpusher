package wallet

import "time"

type Address struct {
	ID int64

	Crypto string
	Text   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
