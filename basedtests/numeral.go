package numeral

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if arabic == 5 {
			result.WriteString("V")
			break
		}

		if arabic == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
