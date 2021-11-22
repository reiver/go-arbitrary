package arbitrary

// PhoneNumber returns an arbitrary phone‐number.
func (arb T) PhoneNumber() string {

	var funcs []func()string = []func()string{
		arb.phonenumber_canada,
		arb.phonenumber_egypt,
	}

	fn := funcs[arb.randomness.Intn(len(funcs))]

	return fn()
}
