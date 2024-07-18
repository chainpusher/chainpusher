package sys

import "sync"

type Map[K, V any] struct {
	container sync.Map
}

func (m *Map[K, V]) Put(key K, value V) {
	m.container.Store(key, value)
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	if value, ok := m.container.Load(key); ok {
		return value.(V), true
	}
	var v V
	return v, false
}

func (m *Map[K, V]) Remove(key K) {
	m.container.Delete(key)
}

func NewMap[K, V any]() *Map[K, V] {
	return &Map[K, V]{}
}
