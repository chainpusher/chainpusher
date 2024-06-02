package monitor

import "github.com/chainpusher/chainpusher/model"

type PlatformWatcherFactory interface {
	CreatePlatformWatcher(platform model.Platform) PlatformWatcher
}
