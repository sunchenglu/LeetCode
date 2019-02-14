package algorithm

func LengthOfLastWord(s string) int {

	ret := 0

	length := len(s)

	if length == 0 {
		return 0
	}

	last := s[0]

	for i := 0; i < length; i++ {

		if s[i] != ' ' {

			if last == ' ' {
				ret = 1
			} else {
				ret += 1
			}

		}

		last = s[i]
	}

	return ret
}
