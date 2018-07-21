package algorithm

func ReverseInteger(x int) int {

	result := 0

	for x/10 != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	result = result*10 + x%10

	if result > (1<<31-1) || result < (-1<<31) {
		return 0
	}

	return result
}
