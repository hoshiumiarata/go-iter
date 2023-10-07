package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestReverse(t *testing.T) {
	it := iter.Ints(0, 5)
	it = iter.Reverse(it)
	slice := iter.ToSlice(it)
	expected := []int{4, 3, 2, 1, 0}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
