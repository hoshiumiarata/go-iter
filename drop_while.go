package iter

type dropWhile[T any] struct {
	it   Iterable[T]
	pred func(T) bool
}

// Next returns the next element in the Iterable and whether it exists.
func (d *dropWhile[T]) Next() (t T, ok bool) {
	for {
		t, ok = d.it.Next()
		if !ok {
			return
		}
		if !d.pred(t) {
			return
		}
	}
}

// DropWhile returns an Iterable that only returns elements from it after the first element that does not satisfy pred.
//
// Example:
//
//	iter.DropWhile(iter.Ints(0, 5), func(i int) bool {
//	  return i < 3
//	})
//
// Produces Iterable[int] with values 3, 4.
func DropWhile[T any](it Iterable[T], pred func(T) bool) Iterable[T] {
	return &dropWhile[T]{
		it:   it,
		pred: pred,
	}
}
