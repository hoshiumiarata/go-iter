package iter

type mapIter[K comparable, V any] struct {
	m map[K]V
}

// Next returns the next element in the Iterable and whether it exists.
func (m *mapIter[K, V]) Next() (kv KV[K, V], ok bool) {
	for k, v := range m.m {
		delete(m.m, k)
		return KV[K, V]{K: k, V: v}, true
	}
	return
}

// FromMap returns an Iterable that iterates over the elements in m.
// m is consumed by this function and should not be used after this function is called.
// The order of iteration is not guaranteed.
//
// Example:
//
//	iter.FromMap(map[string]int{
//	  "a": 1,
//	  "b": 2,
//	})
//
// Produces Iterable[KV[string, int]] with values {"a", 1}, {"b", 2}.
func FromMap[K comparable, V any](m map[K]V) Iterable[KV[K, V]] {
	return &mapIter[K, V]{
		m: m,
	}
}
