package algorithm

import "strings"

func RomanToInt(s string) int {

	result := 0

	single := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	mix := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	for k, v := range mix {
		if strings.Contains(s, k) {
			result = result + v
			s = strings.Replace(s, k, "", 1)
		}
	}

	for i := 0; i < len(s); i++ {
		result = result + single[string(s[i])]
	}

	return result
}
