package algorithm

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	rawNumber := x
	reversNumber := 0

	for x/10 != 0 {
		reversNumber = reversNumber*10 + x%10
		x = x / 10
	}

	reversNumber = reversNumber*10 + x%10

	if reversNumber == rawNumber {
		return true
	} else {
		return false
	}
}
