package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestCatch(t *testing.T) {
	t.Run("Panic", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.Transform(it, func(i int) int {
			if i == 3 {
				panic("panic")
			}
			return i
		})
		it = iter.Catch(it, func(err any) {
			if err != "panic" {
				t.Errorf("expected panic, got %v", err)
			}
		})
		slice := iter.ToSlice(it)
		expected := []int{0, 1, 2}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})

	t.Run("No panic", func(t *testing.T) {
		it := iter.Ints(0, 5)
		it = iter.Transform(it, func(i int) int {
			return i
		})
		it = iter.Catch(it, func(err any) {
			t.Errorf("expected no panic, got %v", err)
		})
		slice := iter.ToSlice(it)
		expected := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
	})
}
