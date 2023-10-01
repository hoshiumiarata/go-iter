package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestWithIndex(t *testing.T) {
	iter := goiter.WithIndex(goiter.FromValues("a", "b", "c"))
	slice := goiter.ToSlice(iter)
	expected := []goiter.MapKV[int, string]{
		{0, "a"},
		{1, "b"},
		{2, "c"},
	}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
