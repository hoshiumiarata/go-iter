package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestEmpty(t *testing.T) {
	it := iter.Empty[any]()
	slice := iter.ToSlice(it)
	if len(slice) != 0 {
		t.Errorf("expected empty slice, got %v", slice)
	}
}
