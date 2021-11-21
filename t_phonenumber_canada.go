package arbitrary

import (
	"fmt"
)

func (arb T) phonenumber_canada() string {

	var areacode string
	{
		var areacodes []string = []string{
			"604",
			"250",
			"778",
			"236",
			"672",
		}

		areacode = areacodes[arb.randomness.Intn(len(areacodes))]
	}

	var localpart1 string
	{
		localpart1 = fmt.Sprint(arb.randomness.Intn(1000))
	}

	var localpart2 string
	{
		localpart2 = fmt.Sprint(arb.randomness.Intn(10000))
	}

	var format string
	{
		var formats []string = []string{
			"%s-%s%s",
			"%s-%s-%s",
			"%s %s %s",
			"%s %s%s",
			"(%s) %s%s",
			"(%s)%s%s",
			"(%s) %s-%s",
			"(%s)%s-%s",
			"(%s) %s %s",
			"(%s)%s %s",
			"%s/%s-%s",
			"+1%s%s%s",
			"+1-%s-%s-%s",
			"+1 %s %s %s",
		}

		format = formats[arb.randomness.Intn(len(formats))]
	}

	var result string
	{
		fmt.Sprintf(format, areacode, localpart1, localpart2)
	}

	return result
}

