package sys_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetCwd(t *testing.T) {
	wd, err := os.Getwd()
	assert.NotNil(t, err)
	assert.NotEqual(t, "", wd)
}
