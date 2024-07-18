package sys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StructKey struct {
	Key string
}

type StructValue struct {
	Value string
}

func TestMap_Get(t *testing.T) {
	m := make(map[int]string)
	m[1] = "1"
	delete(m, 1)
	v := m[1]
	assert.Equal(t, v, "")
}

func TestMap_Put(t *testing.T) {
	im := NewMap[int, string]()
	im.Put(1, "1")

	v, ok := im.Get(1)
	assert.Equal(t, "1", v)
	assert.True(t, ok)

	im.Remove(1)
	v, ok = im.Get(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)

}

func TestMap_PutStructKey(t *testing.T) {
	sm := NewMap[*StructKey, string]()
	key1 := &StructKey{Key: "key1"}
	key2 := &StructKey{Key: "key2"}

	sm.Put(key1, "1")
	sm.Put(key2, "2")

	v, ok := sm.Get(key1)
	assert.True(t, ok)
	v2, ok := sm.Get(key2)
	assert.True(t, ok)
	assert.Equal(t, "1", v)
	assert.Equal(t, "2", v2)
}

func TestMap_PutStructValue(t *testing.T) {
	sm := NewMap[*StructKey, StructValue]()
	key1 := &StructKey{Key: "key1"}
	key2 := &StructKey{Key: "key2"}

	sm.Put(key1, StructValue{Value: "1"})
	sm.Put(key2, StructValue{Value: "2"})

	v, ok := sm.Get(key1)
	assert.True(t, ok)
	v2, ok := sm.Get(key2)
	assert.True(t, ok)
	assert.Equal(t, "1", v.Value)
	assert.Equal(t, "2", v2.Value)
}
