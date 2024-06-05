package model_test

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	// Given
	timestamp := time.Unix(1717565328000, 0)

	// When
	seconds := timestamp.Unix()

	// Then
	if seconds != 1717565328000 {
		t.Errorf("Expected 1717565328000, got %d", seconds)
	}
}
