package cron

import (
	"github.com/chainpusher/chainpusher/module/cron2"
	"github.com/robfig/cron"
)

type Task struct {
	cron *cron.Cron
}

func (t *Task) Start() error {
	cron2.ImediateRun(t.cron)
	return nil
}

func (t *Task) Stop() error {
	return nil
}

func (t *Task) Running() bool {
	return true
}

func NewTask(c *cron.Cron) *Task {
	return &Task{cron: c}
}
