package iter

type reverseIterable[T any] struct {
	slice []T
	i     int
}

// Next returns the next element in the Iterable and whether it exists.
func (it *reverseIterable[T]) Next() (t T, ok bool) {
	if it.i < 0 {
		return t, false
	}
	t = it.slice[it.i]
	it.i--
	return t, true
}

// Reverse returns an Iterable that returns the elements of it in reverse order.
// All elements will be buffered in memory on the call to Reverse.
//
// Example:
//
//	iter.Reverse(iter.Ints(0, 5))
//
// Produces Iterable[int] with values 4, 3, 2, 1, 0.
func Reverse[T any](it Iterable[T]) Iterable[T] {
	slice := ToSlice(it)
	return &reverseIterable[T]{slice: slice, i: len(slice) - 1}
}
