package iter

type stopOnError[T Errorable] struct {
	it       Iterable[T]
	complete bool
}

// Next returns the next element in the Iterable and whether it exists.
func (s *stopOnError[T]) Next() (t T, ok bool) {
	if s.complete {
		return
	}
	t, ok = s.it.Next()
	if !ok {
		return
	}
	if t.Error() != nil {
		s.complete = true
	}
	return
}

// StopOnError returns an Iterable that stops on the first error. It also returns the error as last element.
// Elements of the Iterable must be Errorable (e.g. Result).
//
// Example:
//
//	iter.StopOnError(iter.FromValues(
//		iter.Result[int]{Value: 1},
//		iter.Result[int]{Value: 2},
//		iter.Result[int]{Err: errors.New("error")},
//		iter.Result[int]{Value: 3},
//	))
//
// Produces Iterable[iter.Result[int]] with values Result[int]{Value: 1}, Result[int]{Value: 2}, Result[int]{Err: errors.New("error")}.
func StopOnError[T Errorable](it Iterable[T]) Iterable[T] {
	return &stopOnError[T]{
		it: it,
	}
}
