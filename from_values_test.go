package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestFromValues(t *testing.T) {
	iter := goiter.FromValues(0, 1, 2)
	slice := goiter.ToSlice(iter)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
