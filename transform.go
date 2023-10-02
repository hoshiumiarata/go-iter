package iter

type transform[T, U any] struct {
	it Iterable[T]
	f  func(T) U
}

// Next returns the next element in the Iterable and whether it exists.
func (t *transform[T, U]) Next() (u U, ok bool) {
	tt, ok := t.it.Next()
	if !ok {
		return
	}
	return t.f(tt), true
}

// Transform returns an Iterable that applies f to each element in iter.
//
// Example:
//
//	iter.Transform(iter.Ints(0, 5), func(i int) string {
//	  return fmt.Sprintf("%d", i * 2)
//	})
//
// Produces Iterable[string] with values "0", "2", "4", "6", "8".
func Transform[T, U any](it Iterable[T], f func(T) U) Iterable[U] {
	return &transform[T, U]{
		it: it,
		f:  f,
	}
}
