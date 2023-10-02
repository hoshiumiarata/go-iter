package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestEach(t *testing.T) {
	it := iter.Ints(0, 5)
	var sum int
	iter.Each(it, func(i int) {
		sum += i
	})
	if sum != 10 {
		t.Errorf("expected 10, got %v", sum)
	}
}
