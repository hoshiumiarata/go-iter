package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestAny(t *testing.T) {
	t.Run("Any", func(t *testing.T) {
		it := iter.Ints(0, 5)
		ok := iter.Any(it, func(i int) bool {
			return i > 2
		})
		if !ok {
			t.Errorf("expected true, got %v", ok)
		}
	})

	t.Run("Not any", func(t *testing.T) {
		it := iter.Ints(0, 5)
		ok := iter.Any(it, func(i int) bool {
			return i > 5
		})
		if ok {
			t.Errorf("expected false, got %v", ok)
		}
	})
}
