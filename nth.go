package iter

// Nth returns the nth value from the Iterable and whether it exists.
// Index starts at 0.
//
// Example:
//
//	iter.Nth(iter.Ints(0, 5), 2)
//
// Produces 2, true.
func Nth[T any](it Iterable[T], index int) (t T, ok bool) {
	for i := 0; i <= index; i++ {
		t, ok = it.Next()
		if !ok {
			return
		}
	}
	return
}
