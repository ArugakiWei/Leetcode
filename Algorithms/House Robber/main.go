package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	cur, iL1, iL2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		cur = max(iL1, iL2+nums[i])
		iL2 = iL1
		iL1 = cur
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
