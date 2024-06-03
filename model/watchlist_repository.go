package model

type WatchlistRepository interface {
	In(address []*Transaction) []*Transaction
}
