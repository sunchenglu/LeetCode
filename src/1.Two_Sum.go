package src

func twoSum(nums []int, target int) []int {
	num := len(nums)
	for i:=0; i<num; i++ {
		for j:=i+1; j<num; j++{
			sum := nums[i] + nums[j]
			if sum == target {
				res := [] int{i,j}
				return res
			}
		}
	}

	res := [] int{0, 0}
	return res
}