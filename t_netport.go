package arbitrary

// NetPort returns an arbitrary network port.
func (arb T) NetPort() uint16 {

	var port uint16
	{
		port = uint16(arb.randomness.Intn(65536))

		if 0 == arb.randomness.Intn(5) {
			port = uint16(arb.randomness.Intn(1024))
		}

		if 0 == arb.randomness.Intn(4) {
			port = uint16(arb.randomness.Intn(100))
		}

		if 0 == arb.randomness.Intn(30) {
			port = 80
		}
		if 0 == arb.randomness.Intn(29) {
			port = 79
		}
	}

	return port
}
