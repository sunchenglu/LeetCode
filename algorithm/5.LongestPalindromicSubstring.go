package algorithm


func LongestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	start := 0
	maxLen := 1

	for i := 0; i < len(s); i++ {
		if len(s)-i <= maxLen/2 {
			break
		}
		j := i
		k := i

		for k < len(s)-1 && s[k+1] == s[k] {
			k = k+1
		}

		for k < len(s)-1 && j > 0 && s[k + 1] == s[j - 1] {
			k = k+1
			j = j-1
		}

		len_ := k - j + 1

		if len_ > maxLen {
			start = j
			maxLen = len_
		}

	}

	return s[start: maxLen+start]
}