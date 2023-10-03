package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestAll(t *testing.T) {
	t.Run("All", func(t *testing.T) {
		it := iter.Ints(0, 5)
		ok := iter.All(it, func(i int) bool {
			return i < 5
		})
		if !ok {
			t.Errorf("expected true, got %v", ok)
		}
	})

	t.Run("Not all", func(t *testing.T) {
		it := iter.Ints(0, 5)
		ok := iter.All(it, func(i int) bool {
			return i < 4
		})
		if ok {
			t.Errorf("expected false, got %v", ok)
		}
	})
}
