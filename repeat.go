package iter

type repeat[T any] struct {
	values []T
	i      int
}

// Next returns the next element in the Iterable and whether it exists.
func (r *repeat[T]) Next() (t T, ok bool) {
	t = r.values[r.i]
	r.i++
	if r.i == len(r.values) {
		r.i = 0
	}
	return t, true
}

// Repeat returns an Iterable that repeats the given values indefinitely.
//
// Example:
//
//	iter.Repeat(1, 2, 3)
//
// Produces Iterable[int] with values 1, 2, 3, 1, 2, 3, 1, 2, 3, ...
func Repeat[T any](values ...T) Iterable[T] {
	return &repeat[T]{values: values}
}
