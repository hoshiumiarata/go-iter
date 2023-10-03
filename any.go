package iter

// Any returns true if any element in the iterable satisfies the predicate.
//
// Example:
//
//	iter.Any(iter.Ints(0, 5), func(i int) bool {
//	  return i > 2
//	})
//
// Produces true.
func Any[T any](it Iterable[T], pred func(T) bool) bool {
	for {
		t, ok := it.Next()
		if !ok {
			return false
		}
		if pred(t) {
			return true
		}
	}
}
