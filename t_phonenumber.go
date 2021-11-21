package arbitrary

// PhoneNumber returns an arbitrary phone‐number.
func (arb T) PhoneNumber() string {

	var funcs []func()string = []func()string{
		arb.phonenumber_canada,
	}

	fn := funcs[arb.randomness.Intn(len(funcs))]

	return fn()
}
