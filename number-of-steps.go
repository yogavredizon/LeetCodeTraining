package leetcode

func NumberOfSteps(num int) int {
	step := 0
	for num >= 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}

		step++
	}

	// for i := num; i > 0; i-- {
	// 	step = i
	// 	// if i%2 == 0 {
	// 	// 	i /= 2
	// 	// } else {
	// 	// 	i -= 1
	// 	// }

	// }

	return step
}
