package arbitrary

import (
	"fmt"
)

func (arb T) phonenumber_egypt_3_4(areacode string) string {

	var localpart1 int
	{
		localpart1 = arb.randomness.Intn(1000)
	}

	var localpart2 int
	{
		localpart2 = arb.randomness.Intn(10000)
	}

	var format string
	{
		var formats []string = []string{
			"0%s-%03d-%04d",
			"+20-%s-%03d-%04d",
		}

		format = formats[arb.randomness.Intn(len(formats))]
	}

	var result string
	{
		result = fmt.Sprintf(format, areacode, localpart1, localpart2)
	}

	return result
}

