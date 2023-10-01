package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestFromSlice(t *testing.T) {
	iter := goiter.FromSlice([]int{0, 1, 2})
	slice := goiter.ToSlice(iter)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
