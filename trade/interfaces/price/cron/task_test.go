package cron_test

import (
	"testing"
	"time"

	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	c := cron.New()
	executed := false

	c.AddFunc("*/1 * * * * *", func() {
		executed = true
	})
	c.Start()
	time.Sleep(1 * time.Second)

	assert.True(t, executed)
}

func TestEntries(t *testing.T) {
	c := cron.New()
	executed := false

	c.AddFunc("*/1 * * * * *", func() {
		executed = true
	})
	time.Sleep(1 * time.Second)

	entries := c.Entries()
	entries[0].Job.Run()
	assert.Len(t, entries, 1)
	assert.True(t, executed)
}
