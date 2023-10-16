package iter

// Errorable is an interface for types that can return an error.
type Errorable interface {
	Error() error
}
