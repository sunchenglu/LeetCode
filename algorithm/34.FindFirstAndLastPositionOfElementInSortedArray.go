package algorithm

func SearchRange(nums []int, target int) []int {
	ret := []int{-1, -1}

	length := len(nums)

	if length == 0 {
		return ret
	}

	start := binarySearch(nums, float32(target)-0.5)

	if nums[start] != target {
		return ret
	}

	nums = append(nums, 0)
	end := binarySearch(nums, float32(target)+0.5)

	ret[0] = start
	ret[1] = end - 1
	return ret
}

func binarySearch(nums []int, target float32) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		if target < float32(nums[mid]) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
