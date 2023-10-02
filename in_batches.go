package iter

type inBatches[T any] struct {
	it        Iterable[T]
	batchSize int
}

// Next returns the next element in the Iterable and whether it exists.
func (b *inBatches[T]) Next() (batch []T, ok bool) {
	for i := 0; i < b.batchSize; i++ {
		t, ok := b.it.Next()
		if !ok {
			return batch, len(batch) > 0
		}
		batch = append(batch, t)
	}
	return batch, true
}

// InBatches returns an Iterable that yields batches of size batchSize from it.
// The last batch may be smaller than batchSize.
//
// Example:
//
//	iter.InBatches(iter.Ints(0, 5), 2)
//
// Produces Iterable[[]int] with values [0, 1], [2, 3], [4].
func InBatches[T any](it Iterable[T], batchSize int) Iterable[[]T] {
	return &inBatches[T]{
		it:        it,
		batchSize: batchSize,
	}
}
