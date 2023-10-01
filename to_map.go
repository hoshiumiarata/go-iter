package goiter

// ToMap converts iter to a map.
//
// Example:
//
//	goiter.ToMap(goiter.FromValues(
//	  goiter.MapKV{K: "a", V: 1},
//	  goiter.MapKV{K: "b", V: 2},
//	))
//
// Produces map[string]int with values {"a": 1, "b": 2}.
func ToMap[K comparable, V any](iter Iterable[MapKV[K, V]]) map[K]V {
	m := make(map[K]V)
	for {
		kv, ok := iter.Next()
		if !ok {
			return m
		}
		m[kv.K] = kv.V
	}
}
