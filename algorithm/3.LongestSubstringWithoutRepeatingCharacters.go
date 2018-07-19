package algorithm

func LengthOfLongestSubstring(s string) int {

	length := len(s)

	if length < 2 {
		return length
	}

	start := 0
	longest := 0

	chars := make(map[string]int)

	for i, char := range s {

		num, exist := chars[string(char)]

		if exist && start <= num {
			start = num + 1
		} else {
			if longest > (i + 1 - start) {
				longest = longest
			} else {
				longest = i + 1 - start
			}
		}

		chars[string(char)] = i
	}

	return longest
}
