package algorithm

func MaxArea(height []int) int {

	x := len(height) - 1

	if x <= 0 {
		return 0
	}

	leftIndex := 0
	rightIndex := x

	maxArea := 0

	for x > 0 {

		leftValue := height[leftIndex]
		rightValue := height[rightIndex]
		y := rightValue

		if leftValue < rightValue {
			y = leftValue
			leftIndex = leftIndex + 1
		}else {
			rightIndex = rightIndex -1
		}

		area := x*y

		if area > maxArea {
			maxArea = area
		}

		x = x-1
	}

	return maxArea
}