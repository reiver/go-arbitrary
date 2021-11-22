package arbitrary

import (
	"fmt"
	"strings"
)

func (arb T) phonenumber_egypt_4_4(areacode string) string {

	var localpart1 int
	{
		localpart1 = arb.randomness.Intn(10000)
	}

	var localpart2 int
	{
		localpart2 = arb.randomness.Intn(10000)
	}

	var format string
	{
		var formats []string = []string{
			"0%s-%04d-%04d",
			"+20-%s-%04d-%04d",
		}

		format = formats[arb.randomness.Intn(len(formats))]
	}

	var result string
	{
		result = fmt.Sprintf(format, areacode, localpart1, localpart2)
	}

	{
		if 0 == arb.randomness.Intn(4) {
			result = strings.ReplaceAll(result, "0", "٠")
			result = strings.ReplaceAll(result, "1", "١")
			result = strings.ReplaceAll(result, "2", "٢")
			result = strings.ReplaceAll(result, "3", "٣")
			result = strings.ReplaceAll(result, "4", "٤")
			result = strings.ReplaceAll(result, "5", "٥")
			result = strings.ReplaceAll(result, "6", "٦")
			result = strings.ReplaceAll(result, "7", "٧")
			result = strings.ReplaceAll(result, "8", "٨")
			result = strings.ReplaceAll(result, "9", "٩")
		}
	}

	return result
}
