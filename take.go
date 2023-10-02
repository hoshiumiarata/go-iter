package iter

type take[T any] struct {
	it Iterable[T]
	n  int
}

// Next returns the next element in the Iterable and whether it exists.
func (t *take[T]) Next() (v T, ok bool) {
	if t.n <= 0 {
		return
	}
	t.n--
	return t.it.Next()
}

// Take returns an Iterable that only returns the first n elements from iter.
//
// Example:
//
//	iter.Take(iter.Ints(0, 5), 2)
//
// Produces Iterable[int] with values 0, 1.
func Take[T any](it Iterable[T], n int) Iterable[T] {
	return &take[T]{
		it: it,
		n:  n,
	}
}
