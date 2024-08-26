package arbitrary

import (
	"fmt"
	"strings"
)

// InternetHostName returns an arbitrary Internet hostname.
func (arb T) InternetHostName() string {

	var partchars string = "0123456789abcdefghijklmnopqrstuvwxyz"

	var tld string = tlds[arb.randomness.Intn(len(tlds))]

	var parts []string

	{
		var numparts int = 1 + arb.randomness.Intn(3)

		for i:=0; i<numparts; i++ {

			var part string

			if arb.Bool() {
				var partlen int = 1 + arb.randomness.Intn(63)
				part = arb.String(partlen, partchars)
			} else {
				part = wordsEnglish[arb.randomness.Intn(len(wordsEnglish))]

				if 0 == (arb.randomness.Int() % 10) {
					part += fmt.Sprintf("-%d", arb.randomness.Intn(1000))
				}
			}

			parts = append(parts, part)
		}
	}

	parts = append(parts, tld)


	return strings.Join(parts, ".")

}
