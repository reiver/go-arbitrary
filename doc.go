/*
Package arbitrary provides tools for generating arbitrary data.

So that you can do things such as:

	src := rand.NewSource( time.Now().UTC().UnixNano() )
	
	arb := arbitrary.New(src)
	
	// Generate an arbitrary bool.
	// Might be false, might be true.
	bl := arb.Bool()
	
	// Generate an arbitrary string.
	s := arb.String()
*/
package arbitrary
