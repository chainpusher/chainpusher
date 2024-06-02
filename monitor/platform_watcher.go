package monitor

type PlatformWatcher interface {
	Start()

	Stop()
}
