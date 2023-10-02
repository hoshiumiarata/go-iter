package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestIndexBy(t *testing.T) {
	m := iter.IndexBy(iter.Ints(0, 5), func(i int) int {
		return i % 2
	})
	expected := map[int]int{
		0: 4,
		1: 3,
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("expected %v, got %v", expected, m)
	}
}
