package iter

// ToMap converts iter to a map.
//
// Example:
//
//	iter.ToMap(iter.FromValues(
//	  iter.KV{K: "a", V: 1},
//	  iter.KV{K: "b", V: 2},
//	))
//
// Produces map[string]int with values {"a": 1, "b": 2}.
func ToMap[K comparable, V any](it Iterable[KV[K, V]]) map[K]V {
	m := make(map[K]V)
	for {
		kv, ok := it.Next()
		if !ok {
			return m
		}
		m[kv.K] = kv.V
	}
}
