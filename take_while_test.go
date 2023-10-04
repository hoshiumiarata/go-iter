package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestTakeWhile(t *testing.T) {
	t.Run("Take while less than 3", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.TakeWhile(it, func(i int) bool {
			return i < 3
		})
		slice := iter.ToSlice(it)
		expected := []int{0, 1, 2}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})
	t.Run("Take all", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.TakeWhile(it, func(i int) bool {
			return true
		})
		slice := iter.ToSlice(it)
		expected := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})
}
