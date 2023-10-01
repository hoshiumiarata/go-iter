package goiter

type transform[T, U any] struct {
	iter Iterable[T]
	f    func(T) U
}

// Next returns the next element in the Iterable and whether it exists.
func (t *transform[T, U]) Next() (u U, ok bool) {
	tt, ok := t.iter.Next()
	if !ok {
		return
	}
	return t.f(tt), true
}

// Transform returns an Iterable that applies f to each element in iter.
//
// Example:
//
//	goiter.Transform(goiter.Ints(0, 5), func(i int) string {
//	  return fmt.Sprintf("%d", i * 2)
//	})
//
// Produces Iterable[string] with values "0", "2", "4", "6", "8".
func Transform[T, U any](iter Iterable[T], f func(T) U) Iterable[U] {
	return &transform[T, U]{
		iter: iter,
		f:    f,
	}
}
