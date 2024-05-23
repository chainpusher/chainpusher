package commands

import (
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	command := NewMonitorCommand()
	go command.Execute()

	time.Sleep(20 * time.Second)
}
