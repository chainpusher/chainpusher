package secret_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/secret"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecret(t *testing.T) {
	s, err := secret.NewSecret()
	assert.Nil(t, err)
	assert.NotNil(t, s)
}
