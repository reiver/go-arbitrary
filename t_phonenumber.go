package arbitrary

// PhoneNumber returns an arbitrary phone‐number.
func (arb T) PhoneNumber() string {
	return arb.phonenumber_canada()
}
