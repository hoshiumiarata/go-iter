package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestFilter(t *testing.T) {
	it := iter.Filter(iter.Ints(0, 5), func(i int) bool {
		return i%2 == 0
	})
	slice := iter.ToSlice(it)
	expected := []int{0, 2, 4}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
