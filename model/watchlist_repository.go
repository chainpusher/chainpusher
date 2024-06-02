package model

type WatchlistRepository interface {
	IsOnList(address string) bool
}
