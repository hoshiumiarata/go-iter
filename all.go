package iter

// All returns true if all elements in the iterable satisfy the predicate.
//
// Example:
//
//	iter.All(iter.Ints(0, 5), func(i int) bool {
//	  return i < 5
//	})
//
// Produces true.
func All[T any](it Iterable[T], pred func(T) bool) bool {
	for {
		t, ok := it.Next()
		if !ok {
			return true
		}
		if !pred(t) {
			return false
		}
	}
}
