package sys_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/sys"
)

func TestGetEnvAndLoadDotEnv(t *testing.T) {

	infuraKey, err := sys.GetEnv("INFURA_KEY")
	if err != nil {
		t.Log("Failed to get Infura key: ", err)
	}
	t.Log("Infura key: ", infuraKey)
}
