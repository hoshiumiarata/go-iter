package goiter

type mapIter[K comparable, V any] struct {
	m map[K]V
}

// Next returns the next element in the Iterable and whether it exists.
func (m *mapIter[K, V]) Next() (kv MapKV[K, V], ok bool) {
	for k, v := range m.m {
		delete(m.m, k)
		return MapKV[K, V]{K: k, V: v}, true
	}
	return
}

// FromMap returns an Iterable that iterates over the elements in m.
// m is consumed by this function and should not be used after this function is called.
// The order of iteration is not guaranteed.
//
// Example:
//
//	goiter.FromMap(map[string]int{
//	  "a": 1,
//	  "b": 2,
//	})
//
// Produces Iterable[MapKV[string, int]] with values {"a", 1}, {"b", 2}.
func FromMap[K comparable, V any](m map[K]V) Iterable[MapKV[K, V]] {
	return &mapIter[K, V]{
		m: m,
	}
}
