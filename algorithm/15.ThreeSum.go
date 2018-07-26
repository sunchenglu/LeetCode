package algorithm

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	var res [][]int

	length := len(nums)

	if length < 3 {
		return res
	}

	sort.Ints(nums)

	for i:=0; i<length-2; i++ {

		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		head := i+1
		tail := length - 1

		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				head = head+1
				tail = tail -1
				for (head < tail) && nums[head] == nums[head-1] {
					head = head+1
				}
				for (head < tail) && nums[tail] == nums[tail+1] {
					tail = tail-1
				}
			}else if sum < 0 {
				head = head + 1
			}else {
				tail = tail -1
			}
		}

	}

	return res
}