package iter

// IndexBy returns a map of values in it indexed by key.
// If multiple values have the same key, the last value will be used.
//
// Example:
//
//	iter.IndexBy(iter.Ints(0, 5), func(i int) int {
//	  return i % 2
//	})
//
// Produces map[int]int with values 0: 4, 1: 3.
func IndexBy[T any, K comparable](it Iterable[T], key func(T) K) map[K]T {
	m := make(map[K]T)
	for {
		t, ok := it.Next()
		if !ok {
			return m
		}
		m[key(t)] = t
	}
}
