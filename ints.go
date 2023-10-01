package goiter

type ints struct {
	start, end, current int
}

// Next returns the next int in the Iterable and whether it exists.
func (i *ints) Next() (int, bool) {
	if i.current >= i.end {
		return 0, false
	}
	i.current++
	return i.current - 1, true
}

// Ints returns an Iterable of ints from start to end. The Iterable will not include end. start must be less than end.
//
// Example:
//
//	goiter.Ints(0, 5)
//
// Produces Iterable[int] with values 0, 1, 2, 3, 4.
func Ints(start, end int) Iterable[int] {
	return &ints{
		start:   start,
		end:     end,
		current: start,
	}
}
