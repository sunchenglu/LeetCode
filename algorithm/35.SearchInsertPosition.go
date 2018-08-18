package algorithm

func SearchInsert(nums []int, target int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	low := 0
	high := length - 1

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
