package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestToSlice(t *testing.T) {
	iter := goiter.Ints(0, 5)
	slice := goiter.ToSlice(iter)
	expected := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
