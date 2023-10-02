package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestReduce(t *testing.T) {
	t.Run("Not empty", func(t *testing.T) {
		it := iter.Ints(0, 5)
		sum := iter.Reduce(it, func(acc, current int) int {
			return acc + current
		})
		if sum != 10 {
			t.Errorf("expected %v, got %v", 10, sum)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		it := iter.Empty[int]()
		sum := iter.Reduce(it, func(acc, current int) int {
			return acc + current
		})
		if sum != 0 {
			t.Errorf("expected %v, got %v", 0, sum)
		}
	})
}
