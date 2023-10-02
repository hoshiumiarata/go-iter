package iter

// First returns the first element of an iterable that satisfies a condition.
//
// Example:
//
//	iter.First(iter.Ints(1, 5), func(i int) bool {
//	  return i % 2 == 0
//	})
//
// Produces 2, true.
func First[T any](it Iterable[T], cond func(T) bool) (T, bool) {
	for {
		t, ok := it.Next()
		if !ok {
			return t, false
		}
		if cond(t) {
			return t, true
		}
	}
}
