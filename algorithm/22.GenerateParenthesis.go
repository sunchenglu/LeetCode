package algorithm

func GenerateParenthesis(n int) []string {
	return genParenthesis("", n, n)
}

func genParenthesis(prefix string, nLeft, nRight int) []string {
	if nLeft < 0 || nRight < 0 {
		return nil
	}

	if nLeft == 0 && nRight == 0 {
		return []string{prefix}
	}

	if nLeft > nRight {
		return nil
	}

	if nLeft == nRight {
		return genParenthesis(prefix + "(", nLeft-1, nRight)
	}

	var ret []string
	r := genParenthesis(prefix+"(", nLeft-1, nRight)
	ret = append(ret, r...)
	r = genParenthesis(prefix+")", nLeft, nRight-1)
	ret = append(ret, r...)
	return ret
}
