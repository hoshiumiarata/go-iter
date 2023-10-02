package iter

// Reduce reduces iter to a single value using f.
// If iter is empty, returns the zero value of T.
//
// Example:
//
//	iter.Reduce(iter.Ints(0, 5), func(acc, current int) int {
//	  return acc + current
//	})
//
// Produces 10.
func Reduce[T any](it Iterable[T], f func(acc, current T) T) T {
	acc, ok := it.Next()
	if !ok {
		return acc
	}
	for {
		current, ok := it.Next()
		if !ok {
			return acc
		}
		acc = f(acc, current)
	}
}
