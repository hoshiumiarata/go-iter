package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestZip(t *testing.T) {
	t.Run("First longer", func(t *testing.T) {
		it := iter.Zip(iter.Ints(5, 15), iter.Ints(0, 5))
		slice := iter.ToSlice(it)
		expected := []iter.Pair[int, int]{
			{5, 0},
			{6, 1},
			{7, 2},
			{8, 3},
			{9, 4},
		}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})

	t.Run("Second longer", func(t *testing.T) {
		it := iter.Zip(iter.Ints(0, 5), iter.Ints(5, 15))
		slice := iter.ToSlice(it)
		expected := []iter.Pair[int, int]{
			{0, 5},
			{1, 6},
			{2, 7},
			{3, 8},
			{4, 9},
		}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})
}
