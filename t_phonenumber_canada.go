package arbitrary

import (
	"fmt"
)

func (arb T) phonenumber_canada() string {

	var areacode string
	{
		var areacodes []string = []string{
			// alberta
			"403",
			"780",
			"587",
			"825",
			"368",

			// british columnbia
			"604",
			"250",
			"778",
			"236",
			"672",

			// manitoba
			"204",
			"413",

			// new brunswick
			"506",

			// newfoundland and labrador
			"709",

			// nova scotia and prince edward island
			"782",
			"902",

			// ontario
			"226",
			"249",
			"289",
			"343",
			"365",
			"416",
			"437",
			"519",
			"548",
			"613",
			"647",
			"705",
			"807",
			"905",

			// quebec
			"367",
			"418",
			"581",
			"438",
			"450",
			"514",
			"613",
			"581",
			"819",

			// saskatchewan
			"306",
			"639",

			// yukon, northwest territories and nunavut
			"867",
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
		result = fmt.Sprintf(format, areacode, localpart1, localpart2)
	}

	return result
}

