package goiter

type fromSlice[T any] struct {
	slice []T
	i     int
}

// Next returns the next T in the Iterable and whether it exists.
func (f *fromSlice[T]) Next() (t T, ok bool) {
	if f.i >= len(f.slice) {
		return
	}
	t = f.slice[f.i]
	f.i++
	return t, true
}

// FromSlice returns an Iterable that iterates over the elements in slice.
// slice is consumed by this function and should not be used after this function is called.
//
// Example:
//
//	goiter.FromSlice([]int{1, 2, 3})
//
// Produces Iterable[int] with values 1, 2, 3.
func FromSlice[T any](slice []T) Iterable[T] {
	return &fromSlice[T]{
		slice: slice,
	}
}
