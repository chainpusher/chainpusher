package data

type Pair[K comparable, V comparable] struct {
	Key   K
	Value V
}

func NewPair[K comparable, V comparable](key K, value V) Pair[K, V] {
	return Pair[K, V]{key, value}
}

type Pairs[K comparable, V comparable] struct {
	pairs  []Pair[K, V]
	keys   map[K]V
	values map[V]K
}

func (p Pairs[K, V]) GetValue(key K) V {
	return p.keys[key]
}

func (p Pairs[K, V]) GetKey(key V) K {
	return p.values[key]
}

func (p Pairs[K, V]) Add(pair Pair[K, V]) {
	p.pairs = append(p.pairs, pair)
	p.keys[pair.Key] = pair.Value
	p.values[pair.Value] = pair.Key
}

func NewPairs[K comparable, V comparable](pairs ...Pair[K, V]) Pairs[K, V] {
	keys := make(map[K]V)
	values := make(map[V]K)
	for _, pair := range pairs {
		keys[pair.Key] = pair.Value
		values[pair.Value] = pair.Key
	}
	return Pairs[K, V]{pairs, keys, values}
}
