package goiter

type filter[T any] struct {
	iter Iterable[T]
	cond func(T) bool
}

// Next returns the next element in the Iterable and whether it exists.
func (f *filter[T]) Next() (t T, ok bool) {
	for {
		t, ok = f.iter.Next()
		if !ok {
			return
		}
		if f.cond(t) {
			return
		}
	}
}

// Filter returns an Iterable that only returns elements from iter that satisfy cond.
//
// Example:
//
//	goiter.Filter(goiter.Ints(0, 5), func(i int) bool {
//	  return i % 2 == 0
//	})
//
// Produces Iterable[int] with values 0, 2, 4.
func Filter[T any](iter Iterable[T], cond func(T) bool) Iterable[T] {
	return &filter[T]{
		iter: iter,
		cond: cond,
	}
}
