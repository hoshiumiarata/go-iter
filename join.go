package iter

// Join returns a string that is the result of joining the elements of Iterable with delimiter.
//
// Example:
//
//	iter.Join(iter.FromValues("a", "b", "c"), ", ")
//
// Produces "a, b, c".
func Join(it Iterable[string], delimiter string) string {
	var result string
	for {
		s, ok := it.Next()
		if !ok {
			return result
		}
		if result != "" {
			result += delimiter
		}
		result += s
	}
}
