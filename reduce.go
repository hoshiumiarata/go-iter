package goiter

// Reduce reduces iter to a single value using f.
// If iter is empty, returns the zero value of T.
//
// Example:
//
//	goiter.Reduce(goiter.Ints(0, 5), func(acc, current int) int {
//	  return acc + current
//	})
//
// Produces 10.
func Reduce[T any](iter Iterable[T], f func(acc, current T) T) T {
	acc, ok := iter.Next()
	if !ok {
		return acc
	}
	for {
		current, ok := iter.Next()
		if !ok {
			return acc
		}
		acc = f(acc, current)
	}
}
