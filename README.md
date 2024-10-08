# go-arbitrary

**go-arbitrary** is a Go library which provides tools for generating arbitrary data.

## Example
Here is a typical example:
```
// Generate an arbitrary phone‐number.
phoneNumber := arbitrary.PhoneNumber()

// Generate an arbitrary bool.
// Might be false, might be true.
bl := arbitrary.Bool()

// Generate an arbitrary string.
s := arbitrary.String()
```

Or if you want to provide your own souce of randomness:
```
src := rand.NewSource( time.Now().UTC().UnixNano() )
arb := arbitrary.New(src)

// Generate an arbitrary phone‐number.
phoneNumber := arb.PhoneNumber()

// Generate an arbitrary bool.
// Might be false, might be true.
bl := arb.Bool()

// Generate an arbitrary string.
s := arb.String()
```
## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-arbitrary

[![GoDoc](https://godoc.org/github.com/reiver/go-arbitrary?status.svg)](https://godoc.org/github.com/reiver/go-arbitrary)
