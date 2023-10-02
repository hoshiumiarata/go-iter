package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestWithIndex(t *testing.T) {
	it := iter.WithIndex(iter.FromValues("a", "b", "c"))
	slice := iter.ToSlice(it)
	expected := []iter.MapKV[int, string]{
		{0, "a"},
		{1, "b"},
		{2, "c"},
	}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
