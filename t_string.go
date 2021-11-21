package arbitrary


// String returns an arbitrary string.
//
// String accepts either zero (0), one (1) or two (2) parameters.
// So you can call it like:
//
//	s1 := arb.String()
//
//	s2 := arb.String(10) // Returns an arbitrary string of length 10
//
//	s1 := arb.String(10, "ABCDEF0123456789") // Returns an arbitrary string of length 10,
//	                                         // with characters drawn from "ABCDEF0123456789".
//
//	s1 := arb.String(10, []rune{'a', 'b', 'c'}) // Returns an arbitrary string of length 10,
//	                                            // with characters drawn from []rune{'a', 'b', 'c'}.
//
func (arb T) String(args ...interface{}) string {
	return string(arb.Runes(args...))
}
