package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestTake(t *testing.T) {
	iter := goiter.Ints(0, 5)
	take := goiter.Take(iter, 3)
	slice := goiter.ToSlice(take)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
