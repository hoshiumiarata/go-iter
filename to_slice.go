package iter

// ToSlice converts iter to a slice.
//
// Example:
//
//	iter.ToSlice(iter.FromValues(1, 2, 3))
//
// Produces []int with values 1, 2, 3.
func ToSlice[T any](it Iterable[T]) []T {
	var slice []T
	for {
		t, ok := it.Next()
		if !ok {
			return slice
		}
		slice = append(slice, t)
	}
}
