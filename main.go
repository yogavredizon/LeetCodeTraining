package main

import (
	"fmt"

	lc "github.com/yogavredizon/LeetCodeTraining"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6}
	// // fmt.Println(LinearSearch(arr, 5))
	// fmt.Println(BinarySearch(arr, 6))

	fmt.Println(lc.Add(1))
}

func LinearSearch(arr []int, target int) int {
	step := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			break
		}

		step++
	}

	return step
}

func BinarySearch(arr []int, guess int) (string, int, int) {
	low := 0
	high := len(arr) - 1
	// high 10 {1,2,"3",4,5,//'6',7,8,9}
	warning := ""
	target := 4

	for i := low; i < high; i++ {
		if guess > target {
			high = i - 1
			warning = "Too High"
		} else if guess < target {
			low = i + 1
			warning = "Too Low"
		} else {
			warning = "Correct"
		}
	}

	return warning, high, low
}
