package algorithm

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {

	var result [][]int

	length := len(nums)
	jLength := length - 1

	sort.Ints(nums)

	for i := 0; i < length-3; i++ {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i]*4 > target {
			break
		}

		for j := jLength; i+2 < j; j-- {

			if j != jLength && nums[j] == nums[j+1] {
				continue
			}

			if nums[j]*4 < target {
				break
			}

			sub := target - nums[i] - nums[j]

			low := i + 1
			high := j - 1

			for low < high {

				sum := nums[low] + nums[high]

				// 里面的数大，则高位减一，否则低位加一，否则符合条件
				if sum > sub {
					high -= 1
				} else if sum < sub {
					low += 1
				} else {
					answer := []int{nums[i], nums[low], nums[high], nums[j]}
					result = append(result, answer)

					for low < high && nums[low+1] == nums[low] {
						low += 1
					}

					for low < high && nums[high-1] == nums[high] {
						high -= 1
					}
					high = high - 1
					low = low + 1
				}

			}
		}

	}

	return result
}
