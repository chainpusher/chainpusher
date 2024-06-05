package model

import "time"

type Block struct {
	Height int

	Id string

	CreatedAt time.Time
}
