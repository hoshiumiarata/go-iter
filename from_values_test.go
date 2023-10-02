package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestFromValues(t *testing.T) {
	it := iter.FromValues(0, 1, 2)
	slice := iter.ToSlice(it)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
