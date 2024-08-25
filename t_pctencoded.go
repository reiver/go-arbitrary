package arbitrary

var hexdigs []byte = []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'}

// PctEncoded returns an arbitrary pct-encoded.
func (arb T) PctEncoded() string {
	var buffer [3]byte

	buffer[0] = '%'

	var length int = len(hexdigs)

	buffer[1] = hexdigs[arb.randomness.Intn(length)]
	buffer[2] = hexdigs[arb.randomness.Intn(length)]

	return string(buffer[:])
}
