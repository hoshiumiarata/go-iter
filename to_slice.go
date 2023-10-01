package goiter

// ToSlice converts iter to a slice.
//
// Example:
//
//	goiter.ToSlice(goiter.FromValues(1, 2, 3))
//
// Produces []int with values 1, 2, 3.
func ToSlice[T any](iter Iterable[T]) []T {
	var slice []T
	for {
		t, ok := iter.Next()
		if !ok {
			return slice
		}
		slice = append(slice, t)
	}
}
