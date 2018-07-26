package algorithm

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	compare := strs[0]

	for i := 1; i < len(strs); i++ {

		maxLen := len(compare)

		if len(compare) > len(strs[i]) {
			maxLen = len(strs[i])
		}

		compareNew := ""

		for j := 0; j < maxLen; j++ {
			if compare[j] == strs[i][j] {
				compareNew = compareNew + string(compare[j])
			}else {
				break
			}
		}

		compare = compareNew
	}

	return compare
}
