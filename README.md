# go-arbitrary

**go-arbitrary** is a Go library which provides tools for generating arbitrary data.

## Example
```
src := rand.NewSource( time.Now().UTC().UnixNano() )
arb := arbitrary.New(src)

// Generate an arbitrary bool.
// Might be false, might be true.
bl := arb.Bool()

// Generate an arbitrary string.
s := arb.String()
```
## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-arbitrary

[![GoDoc](https://godoc.org/github.com/reiver/go-arbitrary?status.svg)](https://godoc.org/github.com/reiver/go-arbitrary)
