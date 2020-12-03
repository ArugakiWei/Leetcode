package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{2, 3, 2}))
}

func findErrorNums(nums []int) []int {
	dup, miss := -1, -1
	for _, v := range nums {
		index := abs(v) - 1

		if nums[index] < 0 {
			dup = abs(v)
		} else {
			nums[index] = -nums[index]
		}
	}

	for i, v := range nums {
		if v > 0 {
			miss = i + 1
		}
	}

	return []int{dup, miss}
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
