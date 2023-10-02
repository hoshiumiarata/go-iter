package iter

// FromValues returns an Iterable that iterates over values.
//
// Example:
//
//	iter.FromValues(1, 2, 3)
//
// Produces Iterable[int] with values 1, 2, 3.
func FromValues[T any](values ...T) Iterable[T] {
	return &fromSlice[T]{
		slice: values,
	}
}
