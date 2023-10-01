# goiter

This package package that provides a set of functions for working with strongly typed iterators.
It can be used to work with any type that implements the Iterable interface.
It also can be extended with custom functions.
Check out [documentation](https://pkg.go.dev/github.com/hoshiumiarata/goiter) for more information.

## Example

Example that generates infinite sequence of random numbers and takes first 3 of them.
Then it adds index to each element and converts it to map.

```go
gen := goiter.Generator(func() (int, bool) {
    return rand.Int(), true
})
take := goiter.Take(gen, 3)
withIndex := goiter.WithIndex(take)
m := goiter.ToMap(withIndex)
fmt.Println(m)
// => map[0:8705785567292219898 1:5313896861634354376 2:4299041275511179335]
```