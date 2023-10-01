package goiter

// Iterable is an interface for types that can be iterated over.
type Iterable[T any] interface {
	Next() (T, bool)
}
