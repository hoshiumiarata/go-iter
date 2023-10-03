package iter

type zip[T1 any, T2 any] struct {
	it    Iterable[T1]
	other Iterable[T2]
}

// Next returns the next element in the Iterable and whether it exists.
func (z *zip[T1, T2]) Next() (pair Pair[T1, T2], ok bool) {
	t1, ok := z.it.Next()
	if !ok {
		return
	}
	t2, ok := z.other.Next()
	if !ok {
		return
	}
	return Pair[T1, T2]{t1, t2}, true
}

// Zip returns an Iterable that returns pairs of elements from it and other.
// If it and other are not the same length, the shorter length is used.
//
// Example:
//
//	iter.Zip(iter.Ints(0, 5), iter.Ints(5, 15))
//
// Produces Iterable[Pair[int, int]] with values (0, 5), (1, 6), (2, 7), (3, 8), (4, 9).
func Zip[T1 any, T2 any](it Iterable[T1], other Iterable[T2]) Iterable[Pair[T1, T2]] {
	return &zip[T1, T2]{it, other}
}
