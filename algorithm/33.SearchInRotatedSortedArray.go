package algorithm

import "fmt"

func SearchInRotatedSortedArray(nums []int, target int) int {

	nums = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}

	low := 0
	high := len(nums) - 1

	if high < 0 {
		return -1
	}

	for low < high {

		mid := (low + high) / 2

		fmt.Println(low, high, mid)

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	if nums[low] == target {
		return low
	} else {
		return -1
	}
}
