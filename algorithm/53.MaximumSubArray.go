package algorithm


func MaxSubArray(nums []int) int {

	length := len(nums)
	
	last := nums[0]	// 1
	max := nums[0]

	// 利用处理动态规划
	for i:=1; i<length; i++ {

		if last < 0 {

			if last < nums[i] {
				last = nums[i]
			}

		}else {
			last = last + nums[i]
		}

		if last > max {
			max = last
		}
	}

	return max
}