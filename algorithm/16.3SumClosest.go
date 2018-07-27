package algorithm

func ThreeSumClosest(nums []int, target int) int {

	length := len(nums)

	sum := 0

	if length < 3 {
		for _, v := range nums {
			sum += v
		}
		return sum
	}

	closest := nums[0] + nums[1] + nums[2]
	minDistance := nums[0] + nums[1] + nums[2] - target

	if minDistance < 0 {
		minDistance = -minDistance
	}

	for i := 0; i < length-2; i++ {
		for j := i+1; j < length-1; j++ {
			for k := j+1; k < length; k++ {
				sum = nums[i] + nums[j] + nums[k]
				distance := target - sum
				if distance < 0 {
					distance = -distance
				}

				if distance < minDistance {
					minDistance = distance
					closest = sum
				}
			}
		}

	}

	return closest
}
