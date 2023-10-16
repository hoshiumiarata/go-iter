package iter_test

import (
	"errors"
	"testing"

	"github.com/hoshiumiarata/go-iter"
)

func TestStopOnError(t *testing.T) {
	t.Run("With error", func(t *testing.T) {
		err := errors.New("error")
		it := iter.StopOnError(iter.FromValues(
			iter.Result[int]{Value: 1},
			iter.Result[int]{Value: 2},
			iter.Result[int]{Err: err},
			iter.Result[int]{Value: 3},
		))
		slice := iter.ToSlice(it)
		expected := []iter.Result[int]{
			{Value: 1},
			{Value: 2},
			{Err: err},
		}
		if len(slice) != len(expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
		for i := range slice {
			if slice[i].Value != expected[i].Value || slice[i].Err != expected[i].Err {
				t.Errorf("expected %v, got %v", expected, slice)
			}
		}
	})
	t.Run("Without error", func(t *testing.T) {
		it := iter.StopOnError(iter.FromValues(
			iter.Result[int]{Value: 1},
			iter.Result[int]{Value: 2},
			iter.Result[int]{Value: 3},
		))
		slice := iter.ToSlice(it)
		expected := []iter.Result[int]{
			{Value: 1},
			{Value: 2},
			{Value: 3},
		}
		if len(slice) != len(expected) {
			t.Errorf("expected %v, got %v", expected, slice)
		}
		for i := range slice {
			if slice[i].Value != expected[i].Value || slice[i].Err != expected[i].Err {
				t.Errorf("expected %v, got %v", expected, slice)
			}
		}
	})
}
