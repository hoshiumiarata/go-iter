package iter

// Each calls f for each element in it.
//
// Example:
//
//	iter.Each(iter.Ints(0, 5), func(i int) {
//	  fmt.Println(i)
//	})
//
// Prints: 0, 1, 2, 3, 4
func Each[T any](it Iterable[T], f func(T)) {
	for {
		item, ok := it.Next()
		if !ok {
			break
		}
		f(item)
	}
}
