package algorithm

import "sort"

func CombinationSum(candidates []int, target int) [][]int {

	var res [][]int
	length := len(candidates)

	if length == 0 {
		return res
	}

	sort.Ints(candidates)

	var row []int
	res = doSum(row, candidates, target, res)

	return res
}

func doSum(row []int, candidates []int, target int, res [][]int) [][]int {

	if target == 0 {
		res = append(res, row)
	} else {
		for i, v := range candidates {
			if v <= target {
				tmp := make([]int, 0, len(row))
				tmp = append(tmp, row...)
				tmp = append(tmp, v)
				res = doSum(tmp, candidates[i:], target-v, res)
			} else {
				break
			}
		}
	}

	return res
}
