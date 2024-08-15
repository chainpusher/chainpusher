package context_test

import (
	"github.com/chainpusher/chainpusher/module/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string
}

func TestGet(t *testing.T) {
	scope := context.NewScope()
	context.Add(scope, 1)
	r := context.Get[int](scope, 2)

	assert.Equal(t, 1, r)

	p := &Person{Name: "John"}
	context.Add(scope, p)
	r2 := context.Get[*Person](scope, &Person{})
	assert.Equal(t, "John", r2.Name)
}
