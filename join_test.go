package iter_test

import (
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestJoin(t *testing.T) {
	it := iter.Join(iter.FromValues("a", "b", "c"), ", ")
	if it != "a, b, c" {
		t.Errorf("expected a, b, c, got %v", it)
	}
}
