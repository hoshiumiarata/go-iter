package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestNth(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		it := iter.Ints(1, 5)
		nth, ok := iter.Nth(it, 2)
		if !ok {
			t.Errorf("expected true, got %v", ok)
		}
		if nth != 3 {
			t.Errorf("expected 3, got %v", nth)
		}
	})
	t.Run("Not exists", func(t *testing.T) {
		it := iter.Ints(1, 5)
		_, ok := iter.Nth(it, 5)
		if ok {
			t.Errorf("expected false, got %v", ok)
		}
	})
}
