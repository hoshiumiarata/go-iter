package iter

type takeWhile[T any] struct {
	it   Iterable[T]
	pred func(T) bool
}

// Next returns the next element in the Iterable and whether it exists.
func (t *takeWhile[T]) Next() (value T, ok bool) {
	value, ok = t.it.Next()
	if !ok {
		return
	}
	if !t.pred(value) {
		return value, false
	}
	return
}

// TakeWhile returns an Iterable that only returns elements from it while they satisfy pred.
//
// Example:
//
//	iter.TakeWhile(iter.Ints(0, 5), func(i int) bool {
//	  return i < 3
//	})
//
// Produces Iterable[int] with values 0, 1, 2.
func TakeWhile[T any](it Iterable[T], pred func(T) bool) Iterable[T] {
	return &takeWhile[T]{
		it:   it,
		pred: pred,
	}
}
