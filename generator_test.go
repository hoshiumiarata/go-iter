package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestGenerator(t *testing.T) {
	i := 0
	it := iter.Generator(func() (int, bool) {
		if i >= 5 {
			return 0, false
		}
		i++
		return i, true
	})
	slice := iter.ToSlice(it)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
