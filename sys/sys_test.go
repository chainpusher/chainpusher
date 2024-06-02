package sys_test

import (
	"os"
	"testing"
)

func TestGetCwd(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error("Getwd failed")
	}
	if wd == "" {
		t.Error("Getwd returned empty string")
	}

	t.Log("Working directory is", wd)
}
