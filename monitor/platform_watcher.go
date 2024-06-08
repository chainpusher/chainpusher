package monitor

type PlatformWatcherStatus int

const (
	PlatformWatcherStatusSarting = iota
	PlatformWatcherStatusRunning
	PlatformWatcherStatusStopping
	PlatformWatcherStatusStopped
)

type PlatformWatcher interface {
	Start()

	Stop()
}
