package goiter

type drop[T any] struct {
	iter Iterable[T]
	n    int
}

// Next returns the next element in the Iterable and whether it exists.
func (d *drop[T]) Next() (v T, ok bool) {
	for d.n > 0 {
		if _, ok = d.iter.Next(); !ok {
			return
		}
		d.n--
	}
	return d.iter.Next()
}

// Drop returns an Iterable that drops the first n elements from iter.
//
// Example:
//
//	goiter.Drop(goiter.Ints(0, 5), 2)
//
// Produces Iterable[int] with values 2, 3, 4.
func Drop[T any](iter Iterable[T], n int) Iterable[T] {
	return &drop[T]{
		iter: iter,
		n:    n,
	}
}
