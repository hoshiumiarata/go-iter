package iter

// MapKV is a key-value pair from a map.
type MapKV[K comparable, V any] struct {
	K K
	V V
}
