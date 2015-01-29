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
