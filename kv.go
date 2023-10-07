package iter

// KV is a key-value pair from a map.
type KV[K comparable, V any] struct {
	K K
	V V
}
