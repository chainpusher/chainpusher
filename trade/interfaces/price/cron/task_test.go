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
