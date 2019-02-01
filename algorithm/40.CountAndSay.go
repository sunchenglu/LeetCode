package algorithm

import "strconv"

func CountAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	last := "11"

	for i := 3; i <= n; i++ {

		num := last[0]
		count := 1
		newLast := ""

		length := len(last)

		for j := 0; j < length; j++ {

			cur := last[j]

			if j == length -1 { // 到达最后一次循环
				if cur == num { // 当前数字和上一次数字相等
					count += 1
					newLast += strconv.Itoa(count) + string(num)
				} else { // 当前数字和上一次数字不相等
					newLast += strconv.Itoa(count) + string(num) // 先加上之前的
					newLast += strconv.Itoa(1) + string(cur) // 再加上当前的
				}

			} else { // 没到达最后一次循环
				if cur == num { // 当前数字和上一次数字相等
					if j != 0 { // 第一次不计数
						count += 1
					}
				}else {
					newLast += strconv.Itoa(count) + string(num) // 先加上之前的
					count = 1
					num = cur
				}

			}

		}

		last = newLast
	}

	return last
}
