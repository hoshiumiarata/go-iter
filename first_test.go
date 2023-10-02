package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestFirst(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		it := iter.Ints(1, 5)
		first, ok := iter.First(it, func(i int) bool {
			return i%2 == 0
		})
		if !ok {
			t.Errorf("expected true, got %v", ok)
		}
		if first != 2 {
			t.Errorf("expected 2, got %v", first)
		}
	})

	t.Run("Does not exist", func(t *testing.T) {
		it := iter.Ints(1, 5)
		_, ok := iter.First(it, func(i int) bool {
			return i > 5
		})
		if ok {
			t.Errorf("expected false, got %v", ok)
		}
	})
}
