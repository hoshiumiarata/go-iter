package iter

type empty[T any] struct{}

func (e *empty[T]) Next() (t T, ok bool) {
	return
}

// Empty returns an empty Iterable.
//
// Example:
//
//	iter.Empty[any]()
//
// Produces Iterable[any] with no values.
func Empty[T any]() Iterable[T] {
	return &empty[T]{}
}
