package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestRepeat(t *testing.T) {
	it := iter.Repeat(1, 2, 3)
	take := iter.Take(it, 10)
	slice := iter.ToSlice(take)
	expected := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
