package arbitrary

import (
	"strings"
)

// Password returns an arbitrary password.
func (arb T) Password() string {

	var result string
	{
		result = passwords[arb.randomness.Intn(len(passwords))]
	}

	if 0 == arb.randomness.Intn(7) {

		var storage strings.Builder

		var limit int = 2 + arb.randomness.Intn(6)

		for i:=0; i<limit; i++ {
			if 0 != i {
				storage.WriteRune(' ')
			}

			var s string = words[arb.randomness.Intn(len(words))]

			storage.WriteString(s)
		}

		result = storage.String()
	}

	if 0 == arb.randomness.Intn(97) {
		result = arb.String()
	}

	return result
}
