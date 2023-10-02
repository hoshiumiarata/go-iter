package iter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestToMap(t *testing.T) {
	it := iter.FromValues(iter.MapKV[string, string]{"a", "0"}, iter.MapKV[string, string]{"b", "1"}, iter.MapKV[string, string]{"c", "2"})
	m := iter.ToMap(it)
	expected := map[string]string{
		"a": "0",
		"b": "1",
		"c": "2",
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("expected %v, got %v", expected, m)
	}
}
