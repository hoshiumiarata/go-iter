package iter

type filter[T any] struct {
	it   Iterable[T]
	pred func(T) bool
}

// Next returns the next element in the Iterable and whether it exists.
func (f *filter[T]) Next() (t T, ok bool) {
	for {
		t, ok = f.it.Next()
		if !ok {
			return
		}
		if f.pred(t) {
			return
		}
	}
}

// Filter returns an Iterable that only returns elements from iter that satisfy pred.
//
// Example:
//
//	iter.Filter(iter.Ints(0, 5), func(i int) bool {
//	  return i % 2 == 0
//	})
//
// Produces Iterable[int] with values 0, 2, 4.
func Filter[T any](it Iterable[T], pred func(T) bool) Iterable[T] {
	return &filter[T]{
		it:   it,
		pred: pred,
	}
}
