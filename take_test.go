package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestTake(t *testing.T) {
	it := iter.Ints(0, 5)
	take := iter.Take(it, 3)
	slice := iter.ToSlice(take)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
