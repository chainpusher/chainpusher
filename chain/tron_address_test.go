package chain_test

import (
	"encoding/hex"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
)

func TestAddress(t *testing.T) {
	adr := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	base58decode, err := common.DecodeCheck(adr)
	if err != nil {
		t.Error(err)
	}

	base58encode := common.EncodeCheck(base58decode)
	if base58encode != adr {
		t.Error("base58encode != adr")
	}

	adrHex, err := common.FromHex(adr)
	adr2 := address.HexToAddress(adr)
	adrBytes := adr2.Bytes()

	if err != nil {
		t.Error(err)
	}

	t.Log(adr, adr2, adrBytes, adrHex, base58decode)
}

func TestBase58AddressToHex(t *testing.T) {
	adr := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	base58decode, err := common.DecodeCheck(adr)
	if err != nil {
		t.Error(err)
	}

	// hex string of base58 encoded address
	adr2 := hex.EncodeToString(base58decode)

	if adr2 != "41a614f803b6fd780986a42c78ec9c7f77e6ded13c" {
		t.Error("adr2 != 41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	}
}
