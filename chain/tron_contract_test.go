package chain

import (
	"testing"
)

// Test the abi serialize and deserialize of tron contract
func TestTronConstractAbiSerialize(t *testing.T) {

}

func TestTronNewContract(t *testing.T) {
	contract, err := NewTronContract("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	if err != nil {
		t.Error(err)
	}
	if contract == nil {
		t.Error("contract is nil")
	}
}
