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

		var words []string = languages[arb.randomness.Intn(len(languages))]

		var storage strings.Builder

		var limit int = 2 + arb.randomness.Intn(6)

		for i:=0; i<limit; i++ {
			if 0 != i {
				storage.WriteRune(' ')
			}

			var s string = words[arb.randomness.Intn(len(words))]

			if 0 == arb.randomness.Intn(4) {
				s = strings.ToUpper(s)
			}
			if 0 == arb.randomness.Intn(4) {
				s = strings.ToTitle(s)
			}

			storage.WriteString(s)
		}

		result = storage.String()
	}

	if 0 == arb.randomness.Intn(97) {
		result = arb.String()
	}

	if 0 == arb.randomness.Intn(2) {

		if 0 == arb.randomness.Intn(4) {
			result = strings.Replace(result, " ", ".", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.ReplaceAll(result, " ", "_")
		}
	}

	if 0 == arb.randomness.Intn(2) {

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "A", "4", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "a", "4", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "E", "3", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "e", "3", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "G", "9", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "g", "9", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "I", "1", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "i", "1", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "O", "0", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "o", "0", 1)
		}

		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "S", "$", 1)
		}
		if 0 == arb.randomness.Intn(2) {
			result = strings.Replace(result, "s", "$", 1)
		}
	}

	return result
}
