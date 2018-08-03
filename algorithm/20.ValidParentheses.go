package algorithm

func IsParenthesesValid(s string) bool {

	var stack []byte
	parenthesesMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, exist := parenthesesMap[c]; !exist {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		popVal := stack[len(stack)-1]
		stack = stack[0: len(stack)-1]

		if parenthesesMap[c] != popVal {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}else {
		return false
	}
}
