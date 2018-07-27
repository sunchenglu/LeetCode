package algorithm

func LetterCombinations(digits string) []string {

	letterMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var res []string

	for i := 0; i < len(digits); i++ {

		letters := letterMap[digits[i]]

		if i == 0 {
			for _, v := range letters {
				res = append(res, string([]byte{v}))
			}
			continue
		}

		var newRes []string
		for _, str := range res {
			for _, l := range letters {
				newRes = append(newRes, str + string([]byte{l}))
			}
		}
		res = newRes
	}

	return res
}
