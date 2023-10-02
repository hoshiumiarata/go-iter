package iter

type withIndex[T any] struct {
	it Iterable[T]
	i  int
}

// Next returns the next T in the Iterable and whether it exists.
func (w *withIndex[T]) Next() (kv MapKV[int, T], ok bool) {
	t, ok := w.it.Next()
	if !ok {
		return
	}
	kv.K = w.i
	kv.V = t
	w.i++
	return
}

// WithIndex returns an Iterable that iterates over the elements of iter with their index.
//
// Example:
//
//	iter.WithIndex(iter.FromValues("a", "b", "c"))
//
// Produces Iterable[MapKV[int, string]] with values {0, "a"}, {1, "b"}, {2, "c"}.
func WithIndex[T any](it Iterable[T]) Iterable[MapKV[int, T]] {
	return &withIndex[T]{it, 0}
}
