package goiter_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestFromMap(t *testing.T) {
	iter := goiter.FromMap(map[int]string{
		0: "a",
		1: "b",
		2: "c",
	})

	m := goiter.ToMap(iter)
	expected := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("expected %v, got %v", expected, m)
	}
}
