package goiter

type generator[T any] struct {
	g func() (T, bool)
}

// Next returns the next T in the Iterable and whether it exists.
func (g *generator[T]) Next() (t T, ok bool) {
	return g.g()
}

// Generator returns an Iterable that iterates over the elements returned by g.
//
// Example:
//
//	i := 0
//	goiter.Generator(func() (int, bool) {
//	  if i >= 5 {
//	    return 0, false
//	  }
//	  i++
//	  return i, true
//	})
//
// Produces Iterable[int] with values 1, 2, 3, 4, 5.
func Generator[T any](g func() (T, bool)) Iterable[T] {
	return &generator[T]{
		g: g,
	}
}
