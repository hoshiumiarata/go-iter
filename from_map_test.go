package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestFromMap(t *testing.T) {
	it := iter.FromMap(map[int]string{
		0: "a",
		1: "b",
		2: "c",
	})

	m := iter.ToMap(it)
	expected := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("expected %v, got %v", expected, m)
	}
}
