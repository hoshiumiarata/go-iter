package iter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestTransform(t *testing.T) {
	it := iter.Transform(iter.Ints(0, 5), func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})
	slice := iter.ToSlice(it)
	expected := []string{"0", "2", "4", "6", "8"}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}
