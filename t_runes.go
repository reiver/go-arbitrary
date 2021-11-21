package arbitrary


import "fmt"


// Runes returns an arbitrary []rune.
//
// Runes accepts either zero (0), one (1) or two (2) parameters. (These parameters
// must be specific types and have specific and different meanings.)
//
// So you can call it like:
//
//	rs1 := arb.Runes()
//
//	rs2 := arb.Runes(10) // Returns an arbitrary []rune of length 10
//
//	rs1 := arb.Runes(10, "ABCDEF0123456789") // Returns an arbitrary []rune of length 10,
//	                                         // with characters drawn from "ABCDEF0123456789".
//
//	rs1 := arb.Runes(10, []rune{'a', 'b', 'c'}) // Returns an arbitrary []rune of length 10,
//	                                            // with characters drawn from []rune{'a', 'b', 'c'}.
//
func (arb T) Runes(args ...interface{}) []rune {

	// For the arbitrary []rune, figure out (a) the number of runes to include and (b) which runes
	// to draw from.
	//
	// We store "the number of runes to include" in the variable "num_arbitrary_runes".
	// And we store "which runes to draw from" in the variable "runes".
	//
	// These have default values, but they can be overridden by parameters.
		var num_arbitrary_runes = 13
		var runes []rune = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz")

		switch length := len(args); length {
			case 2:
				switch t := args[1].(type) {
					case string:
						runes = []rune(t)
					case []rune:
						runes = t
					default:
						panic( fmt.Sprintf("Runes only accept a 2nd parameter to type string of type []rune. %T was passed.", args[0]) )
				}
				fallthrough
			case 1:
				switch t := args[0].(type) {
					case int:
						num_arbitrary_runes = t
					default:
						panic( fmt.Sprintf("Runes only accept a 1st parameter to type string of type int. %T was passed.", args[0]) )
				}
			case 0:
				// Nothing here.
			default:
				panic( fmt.Sprintf("Runes only accepts zero (0) or one (1) parameters. %d was passed.", length) )
		}

	// Short circuit.
	//
	// If the "num_arbitrary_runes" is zero (0), then there is
	// no point executing the rest of the code.
		if 1 > num_arbitrary_runes {
/////////////////////// RETURN
			return []rune{}
		}

	// Construct the arbitrary []rune.
		arbitrary_runes := make([]rune, num_arbitrary_runes)

		l := len(runes)
		for i,_ := range arbitrary_runes {
			arbitrary_runes[i] = runes[arb.randomness.Intn(l)]
		}

	// Return.
		return arbitrary_runes
}
