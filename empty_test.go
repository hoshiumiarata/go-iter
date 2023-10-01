package goiter_test

import (
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestEmpty(t *testing.T) {
	iter := goiter.Empty[any]()
	slice := goiter.ToSlice(iter)
	if len(slice) != 0 {
		t.Errorf("expected empty slice, got %v", slice)
	}
}
