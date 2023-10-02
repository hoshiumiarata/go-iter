package iter

type drop[T any] struct {
	it Iterable[T]
	n  int
}

// Next returns the next element in the Iterable and whether it exists.
func (d *drop[T]) Next() (v T, ok bool) {
	for d.n > 0 {
		if _, ok = d.it.Next(); !ok {
			return
		}
		d.n--
	}
	return d.it.Next()
}

// Drop returns an Iterable that drops the first n elements from iter.
//
// Example:
//
//	iter.Drop(iter.Ints(0, 5), 2)
//
// Produces Iterable[int] with values 2, 3, 4.
func Drop[T any](it Iterable[T], n int) Iterable[T] {
	return &drop[T]{
		it: it,
		n:  n,
	}
}
