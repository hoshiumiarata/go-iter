package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestDropWhile(t *testing.T) {
	t.Run("Drop while less than 3", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.DropWhile(it, func(i int) bool {
			return i < 3
		})
		slice := iter.ToSlice(it)
		expected := []int{3, 4}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})

	t.Run("Drop all", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.DropWhile(it, func(i int) bool {
			return true
		})
		slice := iter.ToSlice(it)
		if len(slice) != 0 {
			t.Errorf("expected empty slice, got %v", slice)
		}
	})
}
