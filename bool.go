package arbitrary


// Bool returns an arbitrary bool.
func (arb T) Bool() bool {
	return 0 == arb.randomness.Intn(2)
}
