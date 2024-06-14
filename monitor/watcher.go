package monitor

import (
	"time"

	"github.com/sirupsen/logrus"
)

type NewWatcherOptions struct {
	TimeRequiredForOneCycle int
	IsOneTime               bool
	TickHandler             TickHandler
}

type TickHandler interface {
	Tick()
}

type Tick struct {
}

type Watcher struct {
	TimeRequiredForOneCycle int
	tickHandler             TickHandler
	isOneTime               bool
	TheLastMovement         time.Duration
	nextScheduledTime       time.Time
	stopping                bool
	timer                   *time.Timer
	timerLaunchedAt         time.Time
}

func (w *Watcher) Tick(tick *Tick) {
	w.tickHandler.Tick()
}

func (w *Watcher) Stop() {
	w.stopping = true
}

func (w *Watcher) StartContinuousMovement() {
	for {
		now := time.Now()
		w.nextScheduledTime = w.CalculateNextBlockTime(now)

		var tick *Tick = &Tick{}
		go w.Tick(tick)

		w.StartOneTimeMovement(tick)

		if w.stopping {
			break
		}

		if w.isOneTime {
			break
		}
	}
}

func (w *Watcher) AdvanceTimeTo(time time.Time) {
	if w.timer == nil {
		logrus.Errorf("Timer is nil")
	}

	d := time.Sub(w.timerLaunchedAt)
	logrus.Tracef("Advancing time to: %v, duration: %v", time, d)
	w.timer.Reset(d)
	w.nextScheduledTime = time
}

// func (w *Watcher) StartOneTimeMovement(tick *Tick) {
// 	logrus.Tracef("Starting one time movement, now is : %v", time.Now())
// 	elapsed := 0 * time.Second
// 	timeRequiredForOneCycle := w.GetTimeRequiredForOneCycle()
// 	now := time.Now()
// 	timeToNextBlockGeneration := w.CalculateNextBlockTime(now)

// 	logrus.Tracef("Time to next block generation: %v", timeToNextBlockGeneration)

// 	i := 0
// 	for {
// 		if elapsed >= timeRequiredForOneCycle {
// 			break
// 		}

// 		// if the time to advance to is set and it is before the time to next block generation
// 		if !tick.AdvanceTimeTo.IsZero() && tick.AdvanceTimeTo.Before(timeToNextBlockGeneration) {
// 			logrus.Trace("Quickly advance time to: ", tick.AdvanceTimeTo)
// 			var until time.Duration = time.Until(tick.AdvanceTimeTo)

// 			if until > 0 {
// 				logrus.Tracef("Waiting for %v", until)
// 				time.Sleep(until)
// 				elapsed = elapsed + until
// 				break
// 			}
// 		}

// 		wait := 10 * time.Millisecond
// 		time.Sleep(wait)

// 		elapsed = elapsed + wait
// 		i++
// 	}

// 	w.TheLastMovement = elapsed
// 	logrus.Tracef("Elapsed time: %v, executed %d times", elapsed, i)
// }

func (w *Watcher) StartOneTimeMovement(tick *Tick) {
	logrus.Tracef("Starting one time movement, now is : %v", time.Now())
	w.timer = time.NewTimer(w.GetTimeRequiredForOneCycle())
	w.timerLaunchedAt = time.Now()

	<-w.timer.C
	logrus.Tracef("Starting one time movement, now is : %v", time.Now())
}

func (w *Watcher) GetTimeRequiredForOneCycle() time.Duration {
	return time.Duration(w.TimeRequiredForOneCycle) * time.Second
}

// Calculate the time for the next block generation
func (w *Watcher) CalculateNextBlockTime(blockCreateTime time.Time) time.Time {
	return blockCreateTime.Add(w.GetTimeRequiredForOneCycle())
}

func (w *Watcher) GetCurrentScheduledTime() time.Time {
	return w.nextScheduledTime
}

func NewWatcher(options NewWatcherOptions) *Watcher {
	return &Watcher{
		TimeRequiredForOneCycle: options.TimeRequiredForOneCycle,
		tickHandler:             options.TickHandler,
		isOneTime:               options.IsOneTime,
	}
}
