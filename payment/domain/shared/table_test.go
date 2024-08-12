package shared_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_Values(t *testing.T) {
	m := shared.Table[string, int]{
		"1": 1,
		"2": 2,
	}
	assert.Equal(t, shared.Slice[int]{1, 2}, m.Values())
}
