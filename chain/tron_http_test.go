package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTronHttpGetContron(t *testing.T) {
	client := NewTronHttpClient()
	contract, err := client.GetContract("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")

	if err != nil {
		t.Errorf("Failed to get contract: %v", err)
		return
	}

	assert.Equal(t, true, len(contract.Abi.Entrys) > 0, "they should be equal")
}
