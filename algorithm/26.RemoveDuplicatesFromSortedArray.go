package algorithm

func RemoveDuplicates(nums []int) int {

	length := len(nums)

	if length <= 1 {
		return length
	}

	fast := 0

	for i := 1; i < length; i++ {
		if nums[i] != nums[fast] {
			fast += 1
			nums[fast] = nums[i]
		}
	}

	return fast + 1
}
