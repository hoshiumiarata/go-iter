package iter

type catch[T any] struct {
	it      Iterable[T]
	catcher func(any)
}

// Next returns the next element in the Iterable and whether it exists.
func (it *catch[T]) Next() (t T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			it.catcher(r)
			ok = false
		}
	}()
	return it.it.Next()
}

// Catch returns an Iterable that catches panics from it, calling catcher and stopping iteration.
//
// Example:
//
//	iter.Catch(iter.Generator(func() (int, bool) {
//	  panic("oh no")
//	}), func(err any) {
//	  fmt.Println(err)
//	})
//
// Produces Iterable[int] with no values and prints "oh no".
func Catch[T any](it Iterable[T], catcher func(any)) Iterable[T] {
	return &catch[T]{
		it:      it,
		catcher: catcher,
	}
}
