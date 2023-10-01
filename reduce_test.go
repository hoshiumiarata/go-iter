package goiter_test

import (
	"testing"

	"github.com/hoshiumiarata/goiter"
)

func TestReduce(t *testing.T) {
	t.Run("Not empty", func(t *testing.T) {
		iter := goiter.Ints(0, 5)
		sum := goiter.Reduce(iter, func(acc, current int) int {
			return acc + current
		})
		if sum != 10 {
			t.Errorf("expected %v, got %v", 10, sum)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		iter := goiter.Empty[int]()
		sum := goiter.Reduce(iter, func(acc, current int) int {
			return acc + current
		})
		if sum != 0 {
			t.Errorf("expected %v, got %v", 0, sum)
		}
	})
}
