package monitor

import "time"

type Movement interface {
	BeforeQueryingBlock()

	AfterQueryingBlock()

	WaitTheNextBlockToBeGenerated(watcher *PlatformWatcher)
}

type DefaultMovement struct {
}

func (d *DefaultMovement) BeforeQueryingBlock() {

}

func (d *DefaultMovement) AfterQueryingBlock() {

}

func (d *DefaultMovement) WaitTheNextBlockToBeGenerated(watcher *PlatformWatcher) {
	time.Sleep(watcher.GetTimeForBlockGenerated() * time.Second)
}
