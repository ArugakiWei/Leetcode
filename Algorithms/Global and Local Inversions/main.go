package main

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{0, 1, 3, 2}))
}

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minSuf := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minSuf {
			return false
		}
		minSuf = min(minSuf, nums[i])
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
