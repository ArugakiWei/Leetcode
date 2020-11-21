package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	return max(rob1(nums, 0, len(nums)-2), rob1(nums, 1, len(nums)-1))
}

func rob1(nums []int, start, end int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	cur, iL1, iL2 := 0, 0, 0
	for i := start; i <= end; i++ {
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
