package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestInBatches(t *testing.T) {
	it := iter.InBatches(iter.Ints(0, 5), 2)
	slice := iter.ToSlice(it)
	expected := [][]int{{0, 1}, {2, 3}, {4}}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
