package shared

type Table[K comparable, V any] map[K]V

func (t Table[K, V]) Values() Slice[V] {
	var values Slice[V]
	for _, v := range t {
		values = append(values, v)
	}
	return values
}

func (t Table[K, V]) ForEach(fn func(K, V)) {
	for k, v := range t {
		fn(k, v)
	}
}

func (t Table[K, V]) Remove(k K) {
	delete(t, k)
}

func (t Table[K, V]) RemoveAll() {
	for k := range t {
		t.Remove(k)
	}
}

func (t Table[K, V]) ContainsKey(k K) bool {
	_, ok := t[k]
	return ok
}

func (t Table[K, V]) PutIfAbsent(k K, v V) {
	if !t.ContainsKey(k) {
		t[k] = v
	}
}
