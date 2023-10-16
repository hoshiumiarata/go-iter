package iter

// Result is a generic type for returning a value and an error.
type Result[T any] struct {
	Value T
	Err   error
}

// Error returns the error of the result.
// It implements the Errorable interface.
func (res Result[T]) Error() error {
	return res.Err
}
