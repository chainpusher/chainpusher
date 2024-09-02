package cron2_test

import (
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/module/cron2"
	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"
)

func TestImediateRun(t *testing.T) {
	c := cron.New()
	executed := false

	c.AddFunc("1 * * * * *", func() {
		executed = true
	})

	go cron2.ImediateRun(c)
	time.Sleep(10 * time.Millisecond)
	assert.True(t, executed)
}

func TestImediateStart(t *testing.T) {
	c := cron.New()
	executed := false

	c.AddFunc("1 * * * * *", func() {
		executed = true
	})

	cron2.ImediateStart(c)
	assert.True(t, executed)
}
