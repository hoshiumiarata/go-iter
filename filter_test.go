package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestFilter(t *testing.T) {
	iter := goiter.Filter(goiter.Ints(0, 5), func(i int) bool {
		return i%2 == 0
	})
	slice := goiter.ToSlice(iter)
	expected := []int{0, 2, 4}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
