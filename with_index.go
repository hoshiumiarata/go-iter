package goiter

type withIndex[T any] struct {
	iter Iterable[T]
	i    int
}

// Next returns the next T in the Iterable and whether it exists.
func (w *withIndex[T]) Next() (kv MapKV[int, T], ok bool) {
	t, ok := w.iter.Next()
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
//	goiter.WithIndex(goiter.FromValues("a", "b", "c"))
//
// Produces Iterable[MapKV[int, string]] with values {0, "a"}, {1, "b"}, {2, "c"}.
func WithIndex[T any](iter Iterable[T]) Iterable[MapKV[int, T]] {
	return &withIndex[T]{iter, 0}
}
