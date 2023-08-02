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

	return step
}
