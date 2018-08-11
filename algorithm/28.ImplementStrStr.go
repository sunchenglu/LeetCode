package algorithm

import "fmt"

func StrStr(haystack string, needle string) int {

	if len(haystack) < len(needle) {
		return -1
	}

	if haystack == needle {
		return 0
	}

	for i := 0; i < (len(haystack) - len(needle) + 1); i++ {
		flag := true
		fmt.Println(i)
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag == true {
			return i
		}
	}

	return -1
}
