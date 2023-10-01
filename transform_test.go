package goiter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestTransform(t *testing.T) {
	iter := goiter.Transform(goiter.Ints(0, 5), func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})
	slice := goiter.ToSlice(iter)
	expected := []string{"0", "2", "4", "6", "8"}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
