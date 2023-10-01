package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestToMap(t *testing.T) {
	iter := goiter.FromValues(goiter.MapKV[string, string]{"a", "0"}, goiter.MapKV[string, string]{"b", "1"}, goiter.MapKV[string, string]{"c", "2"})
	m := goiter.ToMap(iter)
	expected := map[string]string{
		"a": "0",
		"b": "1",
		"c": "2",
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("expected %v, got %v", expected, m)
	}
}
