package iter

// GroupBy returns a map of values grouped by the given key.
//
// Example:
//
//	iter.GroupBy(iter.Ints(0, 5), func(i int) int {
//	  return i % 2
//	})
//
// Produces map[int][]int with values 0: {0, 2, 4}, 1: {1, 3}.
func GroupBy[T any, K comparable](it Iterable[T], key func(T) K) map[K][]T {
	m := make(map[K][]T)
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		k := key(t)
		m[k] = append(m[k], t)
	}
	return m
}
