package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestDrop(t *testing.T) {
	t.Run("Less than iter length", func(t *testing.T) {
		iter := goiter.Drop(goiter.Ints(0, 5), 2)
		slice := goiter.ToSlice(iter)
		expected := []int{2, 3, 4}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})
	t.Run("More than iter length", func(t *testing.T) {
		iter := goiter.Drop(goiter.Ints(0, 5), 10)
		slice := goiter.ToSlice(iter)
		if len(slice) != 0 {
			t.Errorf("expected empty slice, got %v", slice)
		}
	})
}
